import 'package:dio/dio.dart';
import 'package:flutter/material.dart';
import 'package:squad_go/main.dart';
import '../core/services/chat_service.dart';
import 'dart:convert';
import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';

class ChatPage extends StatefulWidget {
  final String eventID; // On passe l'ID de l'événement à la page

  const ChatPage({super.key, required this.eventID});

  @override
  State<ChatPage> createState() => _ChatPageState();
}

class _ChatPageState extends State<ChatPage>
    with SingleTickerProviderStateMixin {
  final _controller = TextEditingController();
  final _chatService = ChatService();
  final List<String> _messages = [];
  late AnimationController _animationController;
  late String _currentUserId = '';

  @override
  void initState() {
    super.initState();
    _loadCurrentUser();

    _animationController = AnimationController(
      vsync: this,
      duration: const Duration(milliseconds: 500),
    );
    _animationController.forward();

    _chatService.onMessageReceived = (message) {
      final data = jsonDecode(message);
      final isSelf = data['self'] as bool;
      final content = data['content'] as String;

      setState(() {
        _messages.add(isSelf ? 'Moi: $content' : 'Autre: $content');
      });
    };

    _chatService.connect('ws://localhost:3001/ws');

    _loadMessages(widget.eventID);
  }

  // Fonction pour récupérer l'user_id à partir du token
  Future<void> _loadCurrentUser() async {
    final storage = const FlutterSecureStorage();
    final token = await storage.read(key: dotenv.env['JWT_STORAGE_KEY']!);

    if (token != null) {
      final uri = Uri.http(dotenv.env['API_BASE_URL']!, 'api/users/:userId');

      try {
        final response = await dio.get(uri.toString(),
            options: Options(headers: {
              'Content-Type': 'application/json',
              'Authorization': 'Bearer $token',
            }));

        if (response.statusCode == 200) {
          final data = response.data;
          setState(() {
            _currentUserId = data['id'] ?? '';
          });
        } else if (response.statusCode == 403) {
          debugPrint('Accès refusé : ${response.data}');
          setState(() {
            _currentUserId = '';
          });
        } else {
          debugPrint(
              'Erreur inconnue (${response.statusCode}): ${response.data}');
        }
      } catch (e) {
        debugPrint(
            'Erreur lors de la récupération des informations utilisateur : $e');
      }
    } else {
      debugPrint('Token manquant');
    }
  }

  @override
  void dispose() {
    super.dispose();
    _chatService.closeConnection();
    _animationController.dispose();
  }

  // Fonction pour récupérer les messages de l'événement via l'API
  Future<void> _loadMessages(String eventID) async {
    if (_currentUserId.isEmpty) {
      debugPrint(
          'L\'ID utilisateur n\'est pas initialisé. Impossible de charger les messages.');
      return; // Empêcher toute tentative d'appel si l'utilisateur n'est pas valide
    }
    final uri =
        Uri.http(dotenv.env['API_BASE_URL']!, 'api/message/event/$eventID');
    try {
      final response = await dio.get(uri.toString(),
          options: Options(headers: {
            'Content-Type': 'application/json',
          }));

      if (response.statusCode == 200) {
        final List<dynamic> data = response.data;

        setState(() {
          _messages.clear();
        });

        for (var message in data) {
          final content = message['content'];
          final userId = message['user_id'];
          final userName = message['user_name'];

          final isSelf = userId == _currentUserId;

          setState(() {
            _messages.add(isSelf ? 'Moi: $content' : '$userName: $content');
          });
        }
      } else {
        debugPrint(
            'Erreur lors de la récupération des messages : ${response.statusCode}');
      }
    } catch (e) {
      debugPrint('Erreur lors de la récupération des messages : $e');
    }
  }

  // Fonction pour envoyer un message via WebSocket et l'enregistrer dans la base de données
  void _sendMessage() async {
    final message = _controller.text;
    if (message.isNotEmpty) {
      final messageData = {
        'event_id': widget.eventID,
        'user_id': _currentUserId,
        'content': message,
      };

      // Envoi du message via WebSocket
      _chatService.sendMessage(message);

      final uri = Uri.http(dotenv.env['API_BASE_URL']!, 'api/message');
      try {
        final response = await dio.post(uri.toString(),
            options: Options(headers: {
              'Content-Type': 'application/json',
            }),
            data: jsonEncode(messageData));

        if (response.statusCode == 201) {
          debugPrint('Message envoyé et enregistré');
        } else {
          debugPrint(
              'Erreur lors de l\'enregistrement du message : ${response.statusCode}');
        }
      } catch (e) {
        debugPrint('Erreur lors de l\'enregistrement du message : $e');
      }

      _controller.clear();
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
                color: Colors.grey[200],
                child: ListView.builder(
                  padding: const EdgeInsets.all(8.0),
                  itemCount: _messages.length,
                  itemBuilder: (context, index) {
                    final message = _messages[index];
                    final isSelf = message.startsWith('Moi:');
                    // Récupérer le nom de l'utilisateur qui a envoyé le message
                    final userName = isSelf
                        ? 'Moi' // Si c'est l'utilisateur actuel, afficher "Moi"
                        : message.split(':')[
                            0]; // Sinon, afficher le nom de l'utilisateur récupéré

                    return Align(
                      alignment:
                          isSelf ? Alignment.centerRight : Alignment.centerLeft,
                      child: Column(
                        crossAxisAlignment: isSelf
                            ? CrossAxisAlignment.end
                            : CrossAxisAlignment.start,
                        children: [
                          if (!isSelf)
                            Text(
                              userName,
                              style: const TextStyle(
                                  fontSize: 12, fontWeight: FontWeight.bold),
                            ),
                          Container(
                            padding: const EdgeInsets.symmetric(
                                vertical: 10, horizontal: 15),
                            margin: const EdgeInsets.symmetric(vertical: 5),
                            decoration: BoxDecoration(
                              color: isSelf ? Colors.blue[100] : Colors.white,
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
                              isSelf
                                  ? message.replaceFirst('Moi: ', '')
                                  : message.replaceFirst('$userName: ',
                                      ''), // Afficher le message sans le nom
                              style: const TextStyle(
                                  fontSize: 16, color: Colors.black),
                            ),
                          ),
                        ],
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
                      style: const TextStyle(color: Colors.black),
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
