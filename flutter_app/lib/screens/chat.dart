import 'package:flutter/material.dart';
import 'package:web_socket_channel/web_socket_channel.dart';

class ChatScreen extends StatefulWidget {
  const ChatScreen({super.key});

  @override
  State<ChatScreen> createState() => ChatScreenState();
}

class ChatScreenState extends State<ChatScreen>
    with SingleTickerProviderStateMixin {
  late AnimationController _animationController;
  final TextEditingController _messageController = TextEditingController();
  List<String> messages = []; // Liste pour stocker les messages
  late WebSocketChannel _channel; // Channel WebSocket

  @override
  void initState() {
    super.initState();
    _animationController = AnimationController(
      vsync: this,
      duration: const Duration(milliseconds: 300),
      lowerBound: 0,
      upperBound: 1,
    );
    _animationController.forward();

    // Connecter le client au serveur WebSocket
    _channel = WebSocketChannel.connect(
      Uri.parse('ws://localhost:8080/ws'), // URL de votre serveur WebSocket
    );

    // Ecouter les messages entrants du serveur WebSocket
    _channel.stream.listen((message) {
      setState(() {
        messages.add(message); // Ajoute le message à la liste
      });
    });
  }

  @override
  void dispose() {
    _animationController.dispose();
    _messageController.dispose();
    _channel.sink.close(); // Ferme la connexion WebSocket lors de la fermeture
    super.dispose();
  }

  // Fonction pour envoyer un message
  void _sendMessage() {
    final message = _messageController.text.trim();
    if (message.isNotEmpty) {
      _channel.sink.add(message); // Envoi du message au serveur WebSocket
      setState(() {
        messages.add(message); // Ajoute le message localement
      });
      _messageController.clear(); // Efface le champ de texte après l’envoi
    }
  }

  @override
  Widget build(BuildContext context) {
    return AnimatedBuilder(
      animation: _animationController,
      builder: (context, child) => SlideTransition(
        position: Tween(
          begin: const Offset(0, 0.3),
          end: const Offset(0, 0),
        ).animate(
          CurvedAnimation(
            parent: _animationController,
            curve: Curves.easeInOut,
          ),
        ),
        child: child,
      ),
      child: Scaffold(
        appBar: AppBar(
          title: const Text("Chat de l'événement"),
        ),
        body: Column(
          children: [
            Expanded(
              child: Container(
                color: Colors.grey[
                    200], // Couleur de fond claire pour la zone de messages
                child: ListView.builder(
                  padding: const EdgeInsets.all(8.0),
                  itemCount: messages.length,
                  itemBuilder: (context, index) {
                    return Align(
                      alignment: Alignment.centerLeft,
                      child: Container(
                        padding: const EdgeInsets.symmetric(
                            vertical: 10, horizontal: 15),
                        margin: const EdgeInsets.symmetric(vertical: 5),
                        decoration: BoxDecoration(
                          color: Colors.white,
                          borderRadius: BorderRadius.circular(10),
                          boxShadow: [
                            BoxShadow(
                              color: Colors.grey.withOpacity(0.2),
                              blurRadius: 5,
                              offset: Offset(0, 2),
                            ),
                          ],
                        ),
                        child: Text(
                          messages[index],
                          style: const TextStyle(
                            fontSize: 16,
                            color: Colors.black, // Texte en noir
                          ),
                        ),
                      ),
                    );
                  },
                ),
              ),
            ),
            Padding(
              padding: const EdgeInsets.all(8.0),
              child: Row(
                children: [
                  Expanded(
                    child: TextField(
                      controller: _messageController,
                      style: const TextStyle(
                          color: Colors.black), // Texte visible en noir
                      decoration: InputDecoration(
                        hintText: 'Écrire un message...',
                        hintStyle: const TextStyle(color: Colors.grey),
                        filled: true,
                        fillColor: Colors.white,
                        border: OutlineInputBorder(
                          borderRadius: BorderRadius.circular(20),
                          borderSide: BorderSide.none,
                        ),
                        contentPadding: const EdgeInsets.symmetric(
                            vertical: 10, horizontal: 15),
                      ),
                    ),
                  ),
                  const SizedBox(width: 8),
                  IconButton(
                    icon: const Icon(Icons.send),
                    color: Colors.blue,
                    onPressed: _sendMessage,
                  ),
                ],
              ),
            ),
          ],
        ),
      ),
    );
  }
}
