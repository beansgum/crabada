let Contract = require('web3-eth-contract');
const fs = require('fs');
const retus = require("retus");
const HDWalletProvider = require("@truffle/hdwallet-provider")
const TelegramBot = require('node-telegram-bot-api');

Contract.setProvider(new HDWalletProvider(process.env.BOT_MNEMONIC, "wss://api.avax.network/ext/bc/C/ws"));

// Init contract
let IdleGameContract = "0x82a85407BD612f52577909F4A58bfC6873f14DA8"
let WalletAddress = "0xed3428BcC71d3B0a43Bb50a64ed774bEc57100a8"
var contract = new Contract(JSON.parse(fs.readFileSync('./crabadaabi.json', 'utf8')), IdleGameContract);

const TelegramChatID = -733729902
const POLL_URL = "https://idle-api.crabada.com/public/idle/mines?page=1&status=open&looter_address=" + WalletAddress + "&can_loot=1&limit=8"

// init telegram bot
let bot = new TelegramBot("", { polling: true })
bot.sendMessage(TelegramChatID, 'Active');

bot.on('message', (msg) => {
    console.log(msg)
    const chatId = msg.chat.id;
    if (msg.text == "/ping") {
        bot.sendMessage(chatId, 'pong!');
    }
});

// Setup game teams
let teams = []
teams.push({ ID: 1, Strength: 648 })
teams.push({ ID: 2, Strength: 673 })

var lastPollTime = 0
async function pollGamesAndAttack(team) {
    while (true) {
        try {
            console.log("Polling active games for", team.ID)
            res = await retus(POLL_URL)
            var pollTime = currentTimeSeconds()
            activeGamesResult = JSON.parse(res.body)

            if (activeGamesResult.error_code) {
                console.log("error retrieving active games, code: " + activeGamesResult.error_code + " message: " + activeGamesResult.message)
            } else {
                console.log("Fetched " + activeGamesResult.result.totalRecord + " games")

                activeGames = activeGamesResult.result.data
                for (i = 0; i < activeGames.length; i++) {
                    let game = activeGames[i]

                    if (currentTimeSeconds() - game.start_time > 30) {
                        continue
                    }

                    if (game.start_time < lastPollTime && lastPollTime != 0) {
                        continue
                    }

                    let opponentStrength = game.defense_point
                    console.log("Game ID:" + game.game_id + " Opponent Team ID:" + game.team_id + " Team Strength: " + opponentStrength)

                    if (opponentStrength < team.Strength) {

                        attackSuccessful = false
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

        sleep(0.1)
    }
}

async function attack(game, team) {
    let gameID = game.game_id

    console.log("Attacking:", gameID, "using team:", team.ID)

    var attackError, txHash
    await contract.methods.attack(gameID, team.ID).send({ from: WalletAddress }, function (err, res) {
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
    }, 2100 * 1000);
}

function sleep(seconds) {
    var e = new Date().getTime() + (seconds * 1000);
    while (new Date().getTime() <= e) { }
}

function currentTimeSeconds() {
    return new Date().getTime() / 1000
}

async function main(){
    for(i = 0; i < teams.length; i++) {
        await pollGamesAndAttack(teams[i])
    }
}

main()
