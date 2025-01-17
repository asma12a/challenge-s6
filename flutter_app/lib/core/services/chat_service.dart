import 'package:flutter/material.dart';
import 'package:squad_go/core/utils/constants.dart';
import 'package:web_socket_channel/web_socket_channel.dart';

class ChatService {
  late WebSocketChannel _channel;
  bool isConnected = false;

  Function(String)? onMessageReceived;
  Future<void> connect(String eventID, String userID) async {

    try {
      final url = 'ws://${Constants.apiBaseUrlWs}/ws?event_id=$eventID&user_id=$userID';

      _channel = WebSocketChannel.connect(Uri.parse(url));

      isConnected = true;

      _channel.stream.listen((data) {
        if (onMessageReceived != null) {
          onMessageReceived!(data);
        }
      });

      _channel.sink.done.then((_) {
        isConnected = false;
        debugPrint('Connexion WebSocket fermée');
      });

      debugPrint(
          'Connexion WebSocket réussie avec event_id=$eventID et user_id=$userID');
    } catch (e) {
      debugPrint('Erreur de connexion WebSocket : $e');
    }
  }

  // Méthode pour envoyer un message via la WebSocket
  void sendMessage(String message) async {
    if (isConnected) {
      _channel.sink.add(message);
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
