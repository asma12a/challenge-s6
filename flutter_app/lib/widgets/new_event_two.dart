import 'package:flutter/material.dart';

class NewEventTwo extends StatefulWidget {
  const NewEventTwo({super.key});

  @override
  State<NewEventTwo> createState() => _NewEventTwoState();
}

class _NewEventTwoState extends State<NewEventTwo> {
  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: EdgeInsets.all(20),
      child: Form(
        child: Column(children: [
          TextFormField(
            validator: (value) {
              if (value == null ||
                  value.isEmpty ||
                  value.trim().length <= 1 ||
                  value.trim().length > 50) {
                return 'Doit contenir entre 1 et 50 caractères.';
              }
              return null;
            },
            style: TextStyle(color: Theme.of(context).colorScheme.onSurface),
            maxLength: 50,
            decoration: const InputDecoration(
              border: OutlineInputBorder(),
              icon: Icon(Icons.title),
              label: Text('Nom de l\'équipe'),
            ),
            onSaved: (value) {},
          ),
          TextFormField(
            keyboardType: TextInputType.number,
            decoration: const InputDecoration(
              border: OutlineInputBorder(),
              icon: Icon(Icons.groups),
              label: Text('Nombres de joueurs'),
            ),
          )
        ]),
      ),
    );
  }
}
