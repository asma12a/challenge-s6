import 'package:squad_go/core/models/sport.dart';
import 'package:squad_go/core/services/event_service.dart';
import 'package:squad_go/main.dart';
import 'package:squad_go/screens/tabs.dart';
import 'package:flutter_typeahead/flutter_typeahead.dart';
import 'package:file_picker/file_picker.dart';
import 'package:flutter/material.dart';
import 'package:intl/intl.dart';

class NewEvent extends StatefulWidget {
  const NewEvent({super.key});

  @override
  State<NewEvent> createState() => _NewEventState();
}

class _NewEventState extends State<NewEvent> {
  final _formKey = GlobalKey<FormState>();
  var _enteredName = '';
  var _selectedFile = '';
  String? _selectedDate;
  final TextEditingController _addressController = TextEditingController();
  List<String> _suggestedAddresses = [];
  var _selectedType = '';
  var _selectedSport = '';
  List<Sport> _sports = [];
  final eventService = EventService();

  @override
  void initState() {
    super.initState();
    _initSports();
  }

  Future<void> _initSports() async {
    try {
      final fetchedSports = await eventService.getSports();
      setState(() {
        _sports = fetchedSports;
      });
    } catch (error) {
      ScaffoldMessenger.of(context).clearSnackBars();
      ScaffoldMessenger.of(context).showSnackBar(
        SnackBar(
          content: Text(
            textAlign: TextAlign.center,
            error.toString(),
            style: TextStyle(color: Theme.of(context).colorScheme.onError),
          ),
          backgroundColor: Theme.of(context).colorScheme.error,
        ),
      );
    }
  }

  Future<List<String>> _getAddressSuggestions(String query) async {
    if (query.isEmpty) {
      return [];
    }
    final String apiUrl =
        'https://api-adresse.data.gouv.fr/search/?q=$query&limit=5';
    final response = await dio.get(apiUrl);

    if (response.statusCode == 200) {
      final Map<String, dynamic> data = response.data;

      setState(() {
        _suggestedAddresses = data['features']
            .where((feature) => feature['properties']['label'] != null)
            .map<String>((feature) => feature['properties']['label'] as String)
            .toList();
      });
      return _suggestedAddresses;
    } else {
      throw Exception('Impossible de récupérer les suggestions');
    }
  }

  String? _validateAddress(String? value) {
    if (value == null || value.isEmpty) {
      return 'Le champ adresse ne peut pas être vide.';
    } else if (!_suggestedAddresses.contains(value)) {
      return 'Veuillez sélectionner une adresse correcte.';
    }
    return null;
  }

  void _presentDatePicker() async {
    final now = DateTime.now();
    final firstDate = DateTime(now.year, now.month, now.day);
    final lastDate = DateTime(now.year + 4);

    final pickedDate = await showDatePicker(
      context: context,
      initialDate: now,
      firstDate: firstDate,
      lastDate: lastDate,
    );

    if (pickedDate != null) {
      final formattedDate = DateFormat('yyyy-MM-dd').format(pickedDate);
      setState(() {
        _selectedDate = formattedDate; // _selectedDate devient une chaîne
      });
    }
  }

  void _pickFile() async {
    FilePickerResult? result = await FilePicker.platform.pickFiles();
    if (result != null) {
      PlatformFile file = result.files.first;
      setState(() {
        _selectedFile = file.name;
      });
      debugPrint(file.name);
    } else {
      ScaffoldMessenger.of(context).clearSnackBars();
      ScaffoldMessenger.of(context).showSnackBar(
        SnackBar(
          backgroundColor: Theme.of(context).colorScheme.errorContainer,
          content: Text(
            "Erreur lors de la sélection du document",
            style: TextStyle(
                color: Theme.of(context).colorScheme.onErrorContainer),
            textAlign: TextAlign.center,
          ),
        ),
      );
    }
  }

  void _saveEvent() async {
    if (_formKey.currentState!.validate()) {
      _formKey.currentState!.save();
      if (_selectedDate == null || _selectedDate!.isEmpty) {
        ScaffoldMessenger.of(context).clearSnackBars();
        ScaffoldMessenger.of(context).showSnackBar(
          SnackBar(
            content: Text(
              textAlign: TextAlign.center,
              "Veuillez renseigner une date.",
              style: TextStyle(color: Theme.of(context).colorScheme.onError),
            ),
            backgroundColor: Theme.of(context).colorScheme.error,
          ),
        );
        return;
      }
      try {
        final newEvent = {
          "name": _enteredName,
          "address": _addressController.text,
          "date": _selectedDate!,
          "sport_id": _selectedSport,
          "event_type": _selectedType,
        };
        await eventService.createEvent(newEvent);
        ScaffoldMessenger.of(context).clearSnackBars();
        ScaffoldMessenger.of(context).showSnackBar(
          SnackBar(
            content: Text(
              textAlign: TextAlign.center,
              "L'événement a bien été enregistré.",
              style: TextStyle(color: Theme.of(context).colorScheme.onPrimary),
            ),
            backgroundColor: Theme.of(context).colorScheme.primary,
          ),
        );
        Navigator.of(context).pushReplacement(
          MaterialPageRoute(
            builder: (ctx) => TabsScreen(),
          ),
        );
      } catch (error) {
        ScaffoldMessenger.of(context).clearSnackBars();
        ScaffoldMessenger.of(context).showSnackBar(
          SnackBar(
            content: Text(
              textAlign: TextAlign.center,
              "Serveur indisponible. Veuillez réessayer plus tard.",
              style: TextStyle(color: Theme.of(context).colorScheme.onError),
            ),
            backgroundColor: Theme.of(context).colorScheme.error,
          ),
        );
      }
    }

    @override
    void dispose() {
      // TODO: implement dispose
      _addressController.dispose();
      super.dispose();
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text("Créer un événement"),
      ),
      body: Center(
        child: Column(
          mainAxisSize: MainAxisSize.min,
          children: [
            Padding(
              padding: EdgeInsets.all(20),
              child: Form(
                key: _formKey,
                child: Column(
                  children: [
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
                      style: TextStyle(
                          color: Theme.of(context).colorScheme.onSurface),
                      maxLength: 50,
                      decoration: const InputDecoration(
                        border: OutlineInputBorder(),
                        icon: Icon(Icons.title),
                        label: Text('Nom de l\'événement'),
                      ),
                      onSaved: (value) {
                        _enteredName = value!;
                      },
                    ),
                    SizedBox(height: 20),
                    TypeAheadField<String>(
                      controller: _addressController,
                      // Utilisation du controller
                      builder: (context, controller, focusNode) {
                        return TextFormField(
                          controller: controller,
                          focusNode: focusNode,
                          autofocus: true,
                          style: TextStyle(
                              color: Theme.of(context).colorScheme.onSurface),
                          decoration: InputDecoration(
                            border: OutlineInputBorder(),
                            icon: Icon(Icons.place),
                            labelText: 'Adresse de l\'événement',
                          ),
                          validator: _validateAddress,
                          onSaved: (value) {
                            _addressController.text = value!;
                          },
                        );
                      },
                      suggestionsCallback: (search) =>
                          _getAddressSuggestions(search),
                      itemBuilder: (context, suggestion) {
                        return ListTile(
                          title: Text(suggestion),
                        );
                      },
                      onSelected: (suggestion) {
                        // TODO: Récuperer latitute et longitude (comme dans edit_event.dart)
                        _addressController.text =
                            suggestion; // Mise à jour du champ de texte
                      },
                    ),
                    SizedBox(height: 40),
                    Padding(
                      padding: EdgeInsets.symmetric(horizontal: 2),
                      child: Row(
                        children: [
                          Icon(Icons.calendar_month),
                          SizedBox(width: 15),
                          ElevatedButton(
                            onPressed: _presentDatePicker,
                            child: Text("Sélectionner une date"),
                          ),
                          SizedBox(width: 20),
                          Text(
                            style: TextStyle(
                                color: Theme.of(context).colorScheme.onSurface),
                            _selectedDate == null ? '' : _selectedDate!,
                          ),
                        ],
                      ),
                    ),
                    SizedBox(height: 30),
                    Padding(
                      padding: EdgeInsets.symmetric(horizontal: 2),
                      child: Row(
                        children: [
                          Icon(Icons.sports_sharp),
                          SizedBox(width: 15),
                          SizedBox(
                            width: 200,
                            child: DropdownButtonFormField(
                              validator: (value) {
                                if (value == null) {
                                  return 'Veuillez sélectionner un type.';
                                }
                                return null;
                              },
                              hint: Text(
                                "Type",
                                style: TextStyle(
                                    color: Theme.of(context)
                                        .colorScheme
                                        .onSurface),
                              ),
                              decoration: InputDecoration(
                                border: OutlineInputBorder(),
                              ),
                              items: [
                                DropdownMenuItem(
                                  value: null,
                                  child: Text(
                                    "Type",
                                    style: TextStyle(
                                        color: Theme.of(context)
                                            .colorScheme
                                            .onSurface),
                                  ),
                                ),
                                DropdownMenuItem(
                                  value: "match",
                                  child: Text(
                                    "Match",
                                    style: TextStyle(
                                        color: Theme.of(context)
                                            .colorScheme
                                            .onSurface),
                                  ),
                                ),
                                DropdownMenuItem(
                                  value: "training",
                                  child: Text(
                                    "Training",
                                    style: TextStyle(
                                        color: Theme.of(context)
                                            .colorScheme
                                            .onSurface),
                                  ),
                                )
                              ],
                              onChanged: (value) {
                                _selectedType = value!;
                              },
                            ),
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
                          SizedBox(width: 15),
                          SizedBox(
                            width: 200,
                            child: DropdownButtonFormField(
                              validator: (value) {
                                if (value == null) {
                                  return 'Veuillez sélectionner un sport.';
                                }
                                return null;
                              },
                              hint: Text(
                                "Sport",
                                style: TextStyle(
                                    color: Theme.of(context)
                                        .colorScheme
                                        .onSurface),
                              ),
                              decoration: InputDecoration(
                                border: OutlineInputBorder(),
                              ),
                              items: [
                                DropdownMenuItem<String>(
                                  value: null,
                                  child: Text(
                                    "Sport",
                                    style: TextStyle(
                                      color: Theme.of(context)
                                          .colorScheme
                                          .onSurface,
                                    ),
                                  ),
                                ),
                                ..._sports.map((sport) {
                                  return DropdownMenuItem<String>(
                                    value: sport.id,
                                    child: Text(
                                      sport.name.name,
                                      style: TextStyle(
                                        color: Theme.of(context)
                                            .colorScheme
                                            .onSurface,
                                      ),
                                    ),
                                  );
                                }),
                              ],
                              onChanged: (value) {
                                _selectedSport = value!;
                              },
                            ),
                          ),
                        ],
                      ),
                    ),
                    SizedBox(height: 30),
                    Padding(
                      padding: EdgeInsets.symmetric(horizontal: 2),
                      child: Row(
                        children: [
                          Icon(Icons.image),
                          SizedBox(width: 15),
                          ElevatedButton(
                            onPressed: _pickFile,
                            child: Text("Choisir une image"),
                          ),
                          SizedBox(width: 15),
                          Expanded(
                              child: Text(
                            _selectedFile ?? '',
                            overflow: TextOverflow.ellipsis,
                            style: TextStyle(
                                color: Theme.of(context).colorScheme.onSurface),
                          )),
                        ],
                      ),
                    ),
                    SizedBox(
                      height: 60,
                    ),
                    ElevatedButton(
                        onPressed: _saveEvent, child: Text("Enregistrer"))
                  ],
                ),
              ),
            )
          ],
        ),
      ),
    );
  }
}
