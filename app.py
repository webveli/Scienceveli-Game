from flask import Flask, request, jsonify

app = Flask(__name__)

players = {}

@app.route('/start', methods=['POST'])
def start_game():
    username = request.json.get('username')
    if username in players:
        return jsonify({"message": "Player already exists!"}), 400
    players[username] = {"bitcoin": 0, "farms": 0, "animals": 0}
    return jsonify({"message": f"Welcome {username}!", "stats": players[username]})

@app.route('/action', methods=['POST'])
def perform_action():
    username = request.json.get('username')
    action = request.json.get('action')
    player = players.get(username)

    if not player:
        return jsonify({"message": "Player not found!"}), 404

    if action == "plant":
        player["bitcoin"] += 10
    elif action == "buy_farm":
        if player["bitcoin"] >= 50:
            player["farms"] += 1
            player["bitcoin"] -= 50
        else:
            return jsonify({"message": "Not enough Bitcoin!"}), 400
    elif action == "buy_animals":
        if player["bitcoin"] >= 100:
            player["animals"] += 1
            player["bitcoin"] -= 100
        else:
            return jsonify({"message": "Not enough Bitcoin!"}), 400
    else:
        return jsonify({"message": "Invalid action!"}), 400

    return jsonify({"message": "Action performed!", "stats": player})

if __name__ == '__main__':
    app.run(debug=True)
