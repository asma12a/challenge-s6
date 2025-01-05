import 'package:flutter/material.dart';
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
        debugPrint('Message reçu : $data');

        if (onMessageReceived != null) {
          onMessageReceived!(data);
        }
      });
      debugPrint('Connexion WebSocket réussie');
    } catch (e) {
      debugPrint('Erreur de connexion : $e');
    }
  }

  void sendMessage(String message) {
    if (isConnected) {
      _channel.sink.add(message); // Envoi du message au serveur WebSocket
    } else {
      debugPrint('Erreur : Pas de connexion WebSocket');
    }
  }

  // Fermer la connexion
  void closeConnection() {
    _channel.sink.close();
    isConnected = false;
    debugPrint('Connexion WebSocket fermée');
  }
}
