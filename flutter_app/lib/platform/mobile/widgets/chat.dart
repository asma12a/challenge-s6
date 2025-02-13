import 'package:dio/dio.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:squad_go/core/providers/connectivity_provider.dart';
import 'package:squad_go/core/services/chat_service.dart';
import 'dart:convert';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import 'package:squad_go/core/utils/constants.dart';
import 'package:squad_go/main.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';


const apiBaseUrl = String.fromEnvironment('API_BASE_URL');
const jwtStorageToken = String.fromEnvironment('JWT_STORAGE_KEY');

class ChatPage extends StatefulWidget {
  final String eventID;
  final Color? sportColor;
  final bool isEventFinished;


  const ChatPage({super.key, required this.eventID, required this.isEventFinished, this.sportColor});

  @override
  State<ChatPage> createState() => _ChatPageState();
}

class _ChatPageState extends State<ChatPage>
    with SingleTickerProviderStateMixin {
  final TextEditingController _controller = TextEditingController();
  final ChatService _chatService = ChatService();
  final List<String> _messages = [];
  late AnimationController _animationController;
  String _currentUserId = '';

  final storage = const FlutterSecureStorage();

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
      final userName = data['name'] as String;

      debugPrint("message reçu $data");

      setState(() {
        // Vérification si le message reçu n'est pas déjà dans la liste
        if (!_messages
            .contains(isSelf ? 'Moi: $content' : '$userName: $content')) {
          _messages.add(isSelf ? 'Moi: $content' : '$userName: $content');
        }
      });
    };
  }



  Future<void> _initializeChat() async {
    await _loadCurrentUser();
    if (_currentUserId.isNotEmpty) {
      await _loadMessages(widget.eventID);
      await _chatService.connect(widget.eventID, _currentUserId);
    } else {
      debugPrint("Impossible d'initialiser le chat sans ID utilisateur.");
    }
  }

  Future<void> _loadCurrentUser() async {
    final storage = const FlutterSecureStorage();
    final token = await storage.read(key: jwtStorageToken);

    if (token != null) {
      final uri = '${Constants.apiBaseUrl}/api/auth/me';
      try {
        final response = await Dio().get(uri,
            options: Options(headers: {
              'Content-Type': 'application/json',
              'Authorization': 'Bearer $token',
            }));

        if (response.data['id'] != null) {
          setState(() {
            _currentUserId = response.data['id'];
          });
        }
      } catch (e) {
        debugPrint(
            'Erreur lors de la récupération des informations utilisateur : $e');
      }
    }
  }

  Future<void> _loadMessages(String eventID) async {
    final uri = '${Constants.apiBaseUrl}/api/message/event/$eventID';
    try {
      final storage = const FlutterSecureStorage();
      final token = await storage.read(key: jwtStorageToken);
      final response = await Dio().get(uri,
          options: Options(headers: {'Authorization': 'Bearer $token'}));
      if (response.statusCode == 200) {
        setState(() {
          _messages.clear();
          for (var message in response.data) {
            final content = message['content'];
            final userName = message['user_name'];
            final isSelf = message['created_by'] == _currentUserId;
            _messages.add(isSelf ? 'Moi: $content' : '$userName: $content');
          }
        });
      }
    } catch (e) {
      debugPrint('Erreur lors de la récupération des messages : $e');
    }
  }

  Future<void> _sendMessage() async {
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
          final webSocketMessage = jsonEncode({
            'event_id': widget.eventID,
            'user_id': _currentUserId,
            'content': message,
          });
          debugPrint("Envoi du message via WebSocket : $webSocketMessage");
          await _chatService.sendMessage(webSocketMessage);
        }
      } catch (e) {
        debugPrint('Erreur lors de l\'enregistrement du message : $e');
      }
      _controller.clear();
    }
  }

  @override
  void dispose() {
    _chatService.closeConnection();
    _animationController.dispose();
    _controller.dispose();
    super.dispose();
  }



  @override
  Widget build(BuildContext context) {
    final translate = AppLocalizations.of(context);

    var isOnline = context.watch<ConnectivityState>().isConnected;

    if (widget.isEventFinished) {
      return Center(
        child: Text(
          translate?.chat_event_finished ?? 'Les messages sont désactivés car l’événement est terminé.',
          textAlign: TextAlign.center,
        ),
      );
    }

    if (!isOnline) {
      return Center(
        child: Text(
          'Vous êtes hors ligne. Connectez-vous à Internet pour accéder au chat.',
          style: TextStyle(color: Colors.grey, fontSize: 16),
          textAlign: TextAlign.center,
        ),
      );
    }

    return Column(
      children: [
        Expanded(
          child: Container(
            margin: EdgeInsets.only(top: 16),
            decoration: BoxDecoration(
              borderRadius: BorderRadius.circular(16),
              color: widget.sportColor
                  ?.withValues(
                  alpha: 0.03) ??
                  Theme.of(context)
                      .colorScheme
                      .primary
                      .withValues(alpha: 0.03)
            ),
            child: ListView.builder(
              padding: const EdgeInsets.all(8.0),
              itemCount: _messages.length,
              itemBuilder: (context, index) {
                final message = _messages[index];
                final isSelf = message.startsWith('Moi:');
                final userName = isSelf ? 'Moi' : message.split(':')[0];
                final messageText = message.replaceFirst('$userName: ', '');

                return Align(
                  alignment:
                      isSelf ? Alignment.centerRight : Alignment.centerLeft,
                  child: Column(
                    crossAxisAlignment: isSelf
                        ? CrossAxisAlignment.end
                        : CrossAxisAlignment.start,
                    children: [
                      if (!isSelf)
                        Text(userName,
                            style: const TextStyle(
                                fontSize: 12, fontWeight: FontWeight.bold)),
                      Container(
                        padding: const EdgeInsets.symmetric(
                            vertical: 10, horizontal: 15),
                        margin: const EdgeInsets.symmetric(vertical: 5),
                        decoration: BoxDecoration(
                          color: isSelf ? Colors.blue[100] : Colors.white,
                          borderRadius: BorderRadius.circular(10),
                        ),
                        child: Text(messageText,
                            style: const TextStyle(fontSize: 16)),
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
                  onTapOutside: (event) => FocusScope.of(context).unfocus(),
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
                color: isOnline ? Colors.blue : Colors.grey,
                onPressed: isOnline ? _sendMessage : null,
              ),
            ],
          ),
        ),
      ],
    );
  }
}
