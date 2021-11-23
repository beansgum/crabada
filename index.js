let Contract = require('web3-eth-contract');
let Web3 = require('web3');
let ethers = require("ethers")
const fs = require('fs');
const retus = require("retus");
const HDWalletProvider = require("@truffle/hdwallet-provider")
const TelegramBot = require('node-telegram-bot-api');
const utils = require("./utils")

let provider = new HDWalletProvider(process.env.BOT_MNEMONIC, "wss://api.avax.network/ext/bc/C/ws")
Contract.setProvider(provider);
let web3 = new Web3(provider);

// Init contract
let IdleGameContract = "0x82a85407BD612f52577909F4A58bfC6873f14DA8"
let WalletAddress = "0xed3428BcC71d3B0a43Bb50a64ed774bEc57100a8"
var contract = new Contract(JSON.parse(fs.readFileSync('./crabadaabi.json', 'utf8')), IdleGameContract);

const TelegramChatID = -733729902
const POLL_URL = "https://idle-api.crabada.com/public/idle/mines?page=1&status=open&looter_address=" + WalletAddress + "&can_loot=1&limit=8"
const LOOTS_URL = "https://idle-api.crabada.com/public/idle/mines?looter_address=" + WalletAddress + "&page=1&status=open&limit=8"

// init telegram bot
let bot = new TelegramBot("", { polling: true })
bot.sendMessage(TelegramChatID, 'Active');

bot.on('message', (msg) => {
    const chatId = msg.chat.id;
    if (msg.text == "/ping") {
        bot.sendMessage(chatId, 'pong!');
    } else if (msg.text == "/loots") {
        sendActiveLoots()
    }
});
bot.on("polling_error", console.log);
bot.onText(/\/settle (.+)/, (msg, match) => {
    const chatId = msg.chat.id;
    const resp = match[1];
    if (isNaN(resp)) {
        bot.sendMessage(chatId, "Idiot");
    } else {
        settleGame(resp)
    }
});

bot.onText(/\/attack (.+)/, (msg, match) => {
    const chatId = msg.chat.id;
    const resp = match[1];
    if (isNaN(resp)) {
        bot.sendMessage(chatId, "Idiot");
    } else {
        bot.sendMessage(chatId, "Attacking");
        console.log("Received attack message")
        for(i = 0; i < teams.length; i++) {
            if (teams[i].ID == Number(resp)) {
                pollGamesAndAttack(teams[i])
                return
            }
        }

        bot.sendMessage(chatId, "Team was not found with id: "+ resp);
    }
});

// Setup game teams
let teams = []
teams.push({ ID: 1, Strength: 648 })
teams.push({ ID: 2, Strength: 673 })
teams.push({ ID: 3, Strength: 666 })

var lastPollTime = 0
async function pollGamesAndAttack(team) {
    attackSuccessful = false
    while (!attackSuccessful) {
        try {
            console.log("Polling active games for", team.ID, "strength:", team.Strength)
            res = await retus(POLL_URL)
            var pollTime = utils.currentTimeSeconds()
            activeGamesResult = JSON.parse(res.body)

            activeLoots = await makeRequest(LOOTS_URL)

            if (activeGamesResult.error_code) {
                console.log("error retrieving active games, code: " + activeGamesResult.error_code + " message: " + activeGamesResult.message)
            } else {
                console.log("Fetched " + activeGamesResult.result.totalRecord + " games")

                activeGames = activeGamesResult.result.data
                for (i = 0; i < activeGames.length; i++) {
                    let game = activeGames[i]

                    if (utils.currentTimeSeconds() - game.start_time > 30) {
                        continue
                    }

                    let opponentStrength = game.defense_point
                    console.log("Game ID:" + game.game_id + " Opponent Team ID:" + game.team_id + " Team Strength: " + opponentStrength)

                    if (opponentStrength < team.Strength) {

                        await attack(game, team).then(res => {
                            attackSuccessful = res
                        })

                        console.log("Team attack success:", attackSuccessful)
                        if (attackSuccessful) {
                            console.log("Attack was success, returning")
                            return
                        }
                    }
                }
            }

            lastPollTime = pollTime
            console.log("-------------------------------------------")
        } catch (error) {
            console.log("error fetching games:" + error)
        }

        // Redundant sleep cos why not?
        utils.sleep(0.1)
    }

    console.log("Poll function exiting")
}

async function attack(game, team) {
    let gameID = game.game_id

    console.log("Attacking:", gameID, "using team:", team.ID)

    var attackError, txHash
    await contract.methods.attack(gameID, team.ID).send({ from: WalletAddress, gasPrice: 200000000000}, function (err, res) {
        if (err) {
            console.log("Error sending attack tx:", err)
        } else {
            console.log("Game #" + gameID + "" + "Team #" + team.ID + " attack successful: " + res)
            txHash = res
        }
    }).on("receipt", function (receipt) {
        console.log("Receipt:", receipt)
    }).on("error", function (error) {
        attackError = error
    })

    if (attackError) {
        console.log("attack error:", attackError)
        bot.sendMessage(TelegramChatID, "attack error: " + attackError);
        return false
    }

    prepareTeamForNextAttack(gameID, team, txHash)

    return true
}

function prepareTeamForNextAttack(gameID, team, txHash) {
    bot.sendMessage(TelegramChatID, "⚔️ Game #" + gameID + " Team #" + team.ID + " attack was successful, will reattack in an hour.\nhttps://snowtrace.io/tx/" + txHash);

    setTimeout(() => {
        bot.sendMessage(TelegramChatID, 'Game #' + gameID + " is ready to be settled.");
    }, 3700 * 1000);
}

async function sendActiveLoots() {
    console.log("Fetching active loots")
    try {

        let activeLootsResult = await makeRequest(LOOTS_URL)
        activeLoots = activeLootsResult.data
        if (activeLoots.length > 0) {
            console.log("Fetched " + activeLootsResult.totalRecord + " active loots")

            var sb = "Active loots 💰🤑💰\n------------------------------\n"
            sb += "Game\t\tTeam\t\tReady to settle\n"
            for (i = 0; i < activeLoots.length; i++) {
                let loot = activeLoots[i]
                let ready = utils.currentTimeSeconds() - loot.start_time < 3700 ? "❌" : "✅"
                sb += loot.game_id + "\t\t" + loot.attack_team_id + "\t\t" + ready + "\n"
            }

            bot.sendMessage(TelegramChatID, sb);
        } else {
            bot.sendMessage(TelegramChatID, "No active loots.");
            console.log("No active loots")
        }
    } catch (err) {
        console.log(err)
    }
}

async function findMyLootTeam(gameID) {
    let game = await makeRequest("https://idle-api.crabada.com/public/idle/mine/" + gameID)

    if (game.attack_team_owner.toLowerCase() == WalletAddress.toLowerCase()) {
        return { ID: game.attack_team_id, Strength: game.attack_point }
    }

    throw Error("Team not found for game: " + gameID)
}

async function makeRequest(url) {
    res = await retus(url)

    result = JSON.parse(res.body)

    if (result.error_code) {
        throw Error("error making request, code: " + result.error_code + " message: " + result.message)
    }

    return result.result
}

async function settleGame(gameID) {
    console.log("Settling:", gameID)

    try {

        let team = await findMyLootTeam(gameID)

        var settleError, txHash
        await contract.methods.settleGame(gameID).send({ from: WalletAddress }, function (err, res) {
            if (err) {
                console.log("Error sending attack tx:", err)
            } else {
                console.log("Game #" + gameID + " Team #"+team.ID+" settled successful: " + res)
                txHash = res
            }
        }).on("receipt", function (receipt) {
            console.log("Settle receipt:", receipt)
        }).on("error", function (error) {
            settleError = error
        })

        if (settleError) {
            console.log("settle error:", settleError)
            bot.sendMessage(TelegramChatID, "@x442200 settle error: " + settleError);
            return false
        }

        bot.sendMessage(TelegramChatID, "Game #" + gameID+ " Team #"+team.ID+" has been settled.\nhttps://snowtrace.io/tx/" + txHash);

        //attempt another game
        pollGamesAndAttack(team)
    } catch (error) {
        console.log(error)
        bot.sendMessage(TelegramChatID, "Error restarting game:" + err);
    }
}