import 'package:flutter/material.dart';

class NewEvent extends StatefulWidget {
  const NewEvent({super.key});

  @override
  State<NewEvent> createState() => _NewEventState();
}

class _NewEventState extends State<NewEvent> {
  final _formKey = GlobalKey<FormState>();

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text("Créer un événement"),
      ),
      body: Column(
        children: [
          Padding(
            padding: EdgeInsets.all(20),
            child: Form(
              key: _formKey,
              child: Column(
                children: [
                  TextFormField(
                    style: TextStyle(
                        color: Theme.of(context).colorScheme.onSurface),
                    maxLength: 50,
                    decoration: const InputDecoration(
                      border: OutlineInputBorder(),
                      icon: Icon(Icons.title),
                      label: Text('Nom de l\'événement'),
                    ),
                  ),
                  SizedBox(height: 20),
                  TextFormField(
                    style: TextStyle(
                        color: Theme.of(context).colorScheme.onSurface),
                    maxLength: 100,
                    decoration: const InputDecoration(
                      border: OutlineInputBorder(),
                      icon: Icon(Icons.place),
                      label: Text('Adresse de l\'événement'),
                    ),
                  ),
                  SizedBox(height: 20),
                  Padding(
                    padding: EdgeInsets.symmetric(horizontal: 2),
                    child: Row(
                      children: [
                        Icon(Icons.calendar_month),
                        SizedBox(
                          width: 15,
                        ),
                        ElevatedButton(
                            onPressed: () {},
                            child: Text("Sélectionner une date"))
                      ],
                    ),
                  ),
                  SizedBox(height: 30),
                  Padding(
                    padding: EdgeInsets.symmetric(horizontal: 2),
                    child: Row(
                      children: [
                        Icon(Icons.sports_sharp),
                        SizedBox(
                          width: 15,
                        ),
                        SizedBox(
                          width: 200,
                          child: DropdownButtonFormField(
                              decoration: InputDecoration(
                                border: OutlineInputBorder(),
                              ),
                              items: [
                                DropdownMenuItem(
                                  child: Text(
                                    "Type",
                                    style: TextStyle(
                                        color: Theme.of(context)
                                            .colorScheme
                                            .onSurface),
                                  ),
                                ),
                              ],
                              onChanged: (value) {}),
                        ),
                      ],
                    ),
                  ),
                  SizedBox(height: 30),
                  Padding(
                    padding: EdgeInsets.symmetric(horizontal: 2),
                    child: Row(
                      children: [
                        Icon(Icons.sports_soccer),
                        SizedBox(
                          width: 15,
                        ),
                        SizedBox(
                          width: 200,
                          child: DropdownButtonFormField(
                              decoration: InputDecoration(
                                border: OutlineInputBorder(),
                              ),
                              items: [
                                DropdownMenuItem(
                                  child: Text(
                                    "Sport",
                                    style: TextStyle(
                                        color: Theme.of(context)
                                            .colorScheme
                                            .onSurface),
                                  ),
                                ),
                              ],
                              onChanged: (value) {}),
                        ),
                      ],
                    ),
                  ),
                  SizedBox(height: 30,),
                ],
              ),
            ),
          )
        ],
      ),
    );
  }
}
