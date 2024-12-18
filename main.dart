import 'package:flutter/material.dart';
import 'package:audioplayers/audioplayers.dart'; // مكتبة لتشغيل الأصوات

void main() {
  runApp(BitcoinFarmApp());
}

class BitcoinFarmApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: BitcoinFarmScreen(),
    );
  }
}

class BitcoinFarmScreen extends StatefulWidget {
  @override
  _BitcoinFarmScreenState createState() => _BitcoinFarmScreenState();
}

class _BitcoinFarmScreenState extends State<BitcoinFarmScreen> {
  int bitcoin = 0;
  int farms = 0;
  int animals = 0;
  final AudioPlayer _audioPlayer = AudioPlayer();

  void playSound(String sound) {
    _audioPlayer.play(AssetSource(sound)); // تأكد من وضع ملفات الصوت في مجلد assets
  }

  void plantBitcoin() {
    setState(() {
      bitcoin += 10;
    });
    playSound("plant.mp3");
  }

  void buyFarm() {
    if (bitcoin >= 50) {
      setState(() {
        farms++;
        bitcoin -= 50;
      });
      playSound("buy.mp3");
    }
  }

  void buyAnimals() {
    if (bitcoin >= 100) {
      setState(() {
        animals++;
        bitcoin -= 100;
      });
      playSound("success.mp3");
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: Text('Bitcoin Farm')),
      body: Column(
        mainAxisAlignment: MainAxisAlignment.center,
        children: [
          Text('Bitcoin: $bitcoin', style: TextStyle(fontSize: 24)),
          Text('Farms: $farms', style: TextStyle(fontSize: 24)),
          Text('Animals: $animals', style: TextStyle(fontSize: 24)),
          ElevatedButton(onPressed: plantBitcoin, child: Text('Plant Bitcoin')),
          ElevatedButton(onPressed: buyFarm, child: Text('Buy Farm')),
          ElevatedButton(onPressed: buyAnimals, child: Text('Buy Animals')),
        ],
      ),
    );
  }
}
