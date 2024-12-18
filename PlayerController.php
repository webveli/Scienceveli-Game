<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;

class PlayerController extends Controller
{
    private $players = [];

    public function startGame(Request $request)
    {
        $username = $request->input('username');
        if (isset($this->players[$username])) {
            return response()->json(['message' => 'Player already exists!'], 400);
        }

        $this->players[$username] = ['bitcoin' => 0, 'farms' => 0, 'animals' => 0];
        return response()->json(['message' => 'Game started!', 'stats' => $this->players[$username]]);
    }

    public function action(Request $request)
    {
        $username = $request->input('username');
        $action = $request->input('action');

        if (!isset($this->players[$username])) {
            return response()->json(['message' => 'Player not found!'], 404);
        }

        $player = &$this->players[$username];

        switch ($action) {
            case 'plant':
                $player['bitcoin'] += 10;
                break;
            case 'buyFarm':
                if ($player['bitcoin'] >= 50) {
                    $player['farms']++;
                    $player['bitcoin'] -= 50;
                } else {
                    return response()->json(['message' => 'Not enough Bitcoin!'], 400);
                }
                break;
            case 'buyAnimals':
                if ($player['bitcoin'] >= 100) {
                    $player['animals']++;
                    $player['bitcoin'] -= 100;
                } else {
                    return response()->json(['message' => 'Not enough Bitcoin!'], 400);
                }
                break;
            default:
                return response()->json(['message' => 'Invalid action!'], 400);
        }

        return response()->json(['message' => 'Action successful!', 'stats' => $player]);
    }
}
