import 'package:flutter/material.dart';
import '../core/services/chat_service.dart';

class ChatPage extends StatefulWidget {
  const ChatPage({super.key});

  @override
  State<ChatPage> createState() => _ChatPageState();
}

class _ChatPageState extends State<ChatPage> with SingleTickerProviderStateMixin {
  final _controller = TextEditingController();
  final _chatService = ChatService();
  final List<String> _messages = [];
  late AnimationController _animationController;

  @override
  void initState() {
    super.initState();

    // Initialiser l'AnimationController pour les animations d'entrées
    _animationController = AnimationController(
      vsync: this,
      duration: const Duration(milliseconds: 500),
    );

    _animationController.forward(); // Démarrer l'animation

    // Définir la fonction de callback pour traiter les messages reçus
    _chatService.onMessageReceived = (message) {
      setState(() {
        _messages.add('Autre: $message'); // Ajouter le message reçu à la liste
      });
    };

    _chatService.connect('ws://localhost:3001/ws'); // Remplacez avec l'URL de votre WebSocket
  }

  @override
  void dispose() {
    super.dispose();
    _chatService.closeConnection();
    _animationController.dispose(); // Libérer les ressources de l'animation
  }

  void _sendMessage() {
    final message = _controller.text;
    if (message.isNotEmpty) {
      _chatService.sendMessage(message); // Envoyer le message au serveur
      setState(() {
        _messages.add('Moi: $message'); // Ajouter le message envoyé à la liste
      });
      _controller.clear(); // Effacer le champ de texte après l'envoi
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
          title: const Text('Chat'),
        ),
        body: Column(
          children: [
            Expanded(
              child: Container(
                color: Colors.grey[200], // Couleur de fond claire pour la zone de messages
                child: ListView.builder(
                  padding: const EdgeInsets.all(8.0),
                  itemCount: _messages.length,
                  itemBuilder: (context, index) {
                    return Align(
                      alignment: _messages[index].startsWith('Moi:')
                          ? Alignment.centerRight
                          : Alignment.centerLeft, // Aligner à gauche ou droite
                      child: Container(
                        padding: const EdgeInsets.symmetric(
                            vertical: 10, horizontal: 15),
                        margin: const EdgeInsets.symmetric(vertical: 5),
                        decoration: BoxDecoration(
                          color: _messages[index].startsWith('Moi:')
                              ? Colors.blue[100] // Couleur pour les messages envoyés
                              : Colors.white, // Couleur pour les messages reçus
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
                          _messages[index],
                          style: const TextStyle(
                            fontSize: 16,
                            color: Colors.black,
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
                      controller: _controller,
                      style: const TextStyle(color: Colors.black), // Texte visible en noir
                      decoration: InputDecoration(
                        hintText: 'Entrez un message...',
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
                    onPressed: _sendMessage, // Envoi du message
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
