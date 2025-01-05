import 'package:flutter/material.dart';
import 'package:squad_go/core/services/event_service.dart';
import 'package:go_router/go_router.dart';

class JoinEventScreen extends StatefulWidget {
  const JoinEventScreen({super.key});

  @override
  State<JoinEventScreen> createState() => _JoinEventScreenState();
}

class _JoinEventScreenState extends State<JoinEventScreen> {
  final EventService _eventService = EventService();
  final _formKey = GlobalKey<FormState>();
  var _enteredCode = '';

  void _joinEvent() async {
    if (_formKey.currentState!.validate()) {
      _formKey.currentState!.save();
      try {
        var event = await _eventService.getEventByCode(_enteredCode);
        print(event);
        context.go('/sign-in');
      } catch (error) {
        ScaffoldMessenger.of(context).clearSnackBars();
        ScaffoldMessenger.of(context).showSnackBar(
          SnackBar(
            content: Text(
              textAlign: TextAlign.center,
              "Aucun événement ne correspond à ce code.",
              style: TextStyle(color: Theme.of(context).colorScheme.onPrimary),
            ),
            backgroundColor: Theme.of(context).colorScheme.primary,
          ),
        );
      }
    }
  }

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: EdgeInsets.all(20),
      child: Column(
        children: [
          Title(
            color: Colors.red,
            child: Text(
              "Rejoinde un événement",
              style: TextStyle(fontSize: 22),
            ),
          ),
          SizedBox(height: 45),
          Form(
            key: _formKey,
            child: TextFormField(
              validator: (value) {
                if (value == null || value.isEmpty) {
                  return 'Veuillez saisir une valeur.';
                }
                return null;
              },
              keyboardType: TextInputType.number,
              decoration: const InputDecoration(
                border: OutlineInputBorder(),
                icon: Icon(Icons.qr_code),
                label: Text('Saisir le code de l\'événement'),
              ),
              onSaved: (value) {
                _enteredCode = value!;
              },
            ),
          ),
          SizedBox(height: 30),
          ElevatedButton(
            onPressed: _joinEvent,
            child: Text("Rejoindre"),
          )
        ],
      ),
    );
  }
}
