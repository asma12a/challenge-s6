import 'package:dio/dio.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:squad_go/core/exceptions/app_exception.dart';
import 'package:squad_go/core/providers/connectivity_provider.dart';
import 'package:squad_go/main.dart';
import '../../../core/services/chat_service.dart';
import 'dart:convert';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import 'package:squad_go/core/utils/constants.dart';

const apiBaseUrl = String.fromEnvironment('API_BASE_URL');
const jwtStorageToken = String.fromEnvironment('JWT_STORAGE_KEY');

class ChatPage extends StatefulWidget {
  final String eventID;

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
    _initializeChat();

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

    const urlChat = 'ws://${Constants.apiBaseUrl}/ws';
    _chatService.connect(urlChat);
  }

  Future<void> _initializeChat() async {
    await _loadCurrentUser();
    if (_currentUserId.isNotEmpty) {
      await _loadMessages(widget.eventID);
    } else {
      debugPrint('Impossible d\'initialiser le chat sans ID utilisateur.');
    }
  }

  Future<void> _loadCurrentUser() async {
    final storage = const FlutterSecureStorage();
    final token = await storage.read(key: jwtStorageToken);

    if (token != null) {
      final uri = '${Constants.apiBaseUrl}/api/auth/me';
      try {
        final response = await dio.get(uri.toString(),
            options: Options(headers: {
              'Content-Type': 'application/json',
              'Authorization': 'Bearer $token',
            }));

        if (response.data['id'] != null && response.data['id'].isNotEmpty) {
          setState(() {
            _currentUserId = response.data['id'] ?? '';
          });
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

  Future<void> _loadMessages(String eventID) async {
    final uri = '${Constants.apiBaseUrl}/api/message/event/$eventID';

    try {
      final storage = const FlutterSecureStorage();

      final token = await storage.read(key: jwtStorageToken);

      final response = await dio.get(uri,
          options: Options(headers: {
            'Authorization': 'Bearer $token',
            'Cache-Control': 'no-cache',
          }));

      final List<dynamic> data = response.data;

      setState(() {
        _messages.clear();
      });

      for (var message in data) {
        final content = message['content'];
        final userId = message['created_by'];
        final userName = message['user_name'];

        final isSelf = userId == _currentUserId;

        setState(() {
          _messages.add(isSelf ? 'Moi: $content' : '$userName: $content');
        });
      }
    } catch (e) {
      debugPrint('Erreur lors de la récupération des messages : $e');
    }
  }

  void _sendMessage() async {
    final message = _controller.text;
    if (message.isNotEmpty) {
      final messageData = {
        'event_id': widget.eventID,
        'user_id': _currentUserId,
        'content': message,
      };

      final uri = '${Constants.apiBaseUrl}/api/message';

      try {
        final storage = const FlutterSecureStorage();
        final token = await storage.read(key: jwtStorageToken);

        await dio.post(uri,
            options: Options(headers: {
              'Content-Type': 'application/json',
              'Authorization': 'Bearer $token',
            }),
            data: jsonEncode(messageData));

        debugPrint('Message envoyé et enregistré');

        setState(() {
          _messages.add('Moi: $message');
        });

        if (_chatService.isConnected) {
          _chatService.sendMessage(message);
        }
      } catch (e) {
        debugPrint('Erreur lors de l\'enregistrement du message : $e');
      }

      _controller.clear();
    }
  }

  @override
  Widget build(BuildContext context) {
    var isOnline = context.watch<ConnectivityState>().isConnected;

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
          automaticallyImplyLeading: false,
          title: const Text('Chat'),
        ),
        body: Column(
          children: [
            Expanded(
              child: Container(
                color: Theme.of(context).colorScheme.primary.withValues(
                      alpha: (Theme.of(context).colorScheme.primary.a * 0.03),
                    ),
                child: ListView.builder(
                  padding: const EdgeInsets.all(8.0),
                  itemCount: _messages.length,
                  itemBuilder: (context, index) {
                    final message = _messages[index];
                    final isSelf = message.startsWith('Moi:');
                    final userName = isSelf ? 'Moi' : message.split(':')[0];

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
                  // TODO: when offline show a snackbar with a message and disable the send button
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
