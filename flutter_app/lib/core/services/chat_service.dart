import 'dart:convert';
import 'package:web_socket_channel/web_socket_channel.dart';

class ChatService {
  late WebSocketChannel _channel;
  bool isConnected = false;

  Function(String)? onMessageReceived;

  // Connexion WebSocket
  Future<void> connect(String url) async {
    try {
      _channel = WebSocketChannel.connect(Uri.parse(url));
      isConnected = true;

      _channel.stream.listen((data) {
        print('Message reçu : $data');

        if (onMessageReceived != null) {
          onMessageReceived!(data);
        }
      });
      print('Connexion WebSocket réussie');
    } catch (e) {
      print('Erreur de connexion : $e');
    }
  }

  void sendMessage(String message) {
    if (isConnected) {
      _channel.sink.add(message); // Envoi du message au serveur WebSocket
    } else {
      print('Erreur : Pas de connexion WebSocket');
    }
  }

  // Fermer la connexion
  void closeConnection() {
    _channel.sink.close();
    isConnected = false;
    print('Connexion WebSocket fermée');
  }
}
