const express = require('express');
const app = express();
app.use(express.json());

let players = {};

app.post('/api/start', (req, res) => {
    const { username } = req.body;
    if (!players[username]) {
        players[username] = { bitcoin: 0, farms: 0, animals: 0 };
        res.json({ message: `Player ${username} created!`, stats: players[username] });
    } else {
        res.status(400).json({ message: 'Player already exists!' });
    }
});

app.post('/api/action', (req, res) => {
    const { username, action } = req.body;
    const player = players[username];

    if (!player) return res.status(404).json({ message: 'Player not found!' });

    switch (action) {
        case 'plant':
            player.bitcoin += 10;
            break;
        case 'buyFarm':
            if (player.bitcoin >= 50) {
                player.farms++;
                player.bitcoin -= 50;
            } else {
                return res.status(400).json({ message: 'Not enough Bitcoin!' });
            }
            break;
        case 'buyAnimals':
            if (player.bitcoin >= 100) {
                player.animals++;
                player.bitcoin -= 100;
            } else {
                return res.status(400).json({ message: 'Not enough Bitcoin!' });
            }
            break;
        default:
            return res.status(400).json({ message: 'Invalid action!' });
    }

    res.json({ message: 'Action successful!', stats: player });
});

app.listen(3000, () => console.log('Server running on port 3000'));
