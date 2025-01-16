import 'package:squad_go/core/exceptions/app_exception.dart';
import 'package:squad_go/core/models/sport.dart';
import 'package:squad_go/core/services/event_service.dart';
import 'package:flutter_typeahead/flutter_typeahead.dart';
import 'package:flutter/material.dart';
import 'package:intl/intl.dart';
import 'package:squad_go/main.dart';
import 'package:squad_go/platform/mobile/screens/tabs.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';
import 'package:go_router/go_router.dart';

class NewEvent extends StatefulWidget {
  const NewEvent({super.key});

  @override
  State<NewEvent> createState() => _NewEventState();
}

class _NewEventState extends State<NewEvent> {
  final _formKey = GlobalKey<FormState>();
  var _enteredName = '';
  String? _selectedDate;
  final TextEditingController _addressController = TextEditingController();
  List<dynamic> _suggestedAddresses = [];
  var _selectedType = '';
  var _selectedSport = '';
  List<Sport> _sports = [];
  final eventService = EventService();
  var _isPublic = true;
  double _latitude = 46.603354;
  double _longitude = 1.888334;

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
    try {
      final String apiUrl =
          'https://api-adresse.data.gouv.fr/search/?q=$query&limit=5';
      final response = await dio.get(apiUrl);

      final Map<String, dynamic> data = response.data;

      setState(() {
        _suggestedAddresses = data['features'] as List;
      });
      final List<String> addresses = data['features']
          .where((feature) => feature['properties']['label'] != null)
          .map<String>((feature) => feature['properties']['label'] as String)
          .toList();
      return addresses;
    } catch (e) {
      return [];
    }
  }

  String? _validateAddress(String? value) {
    final translate = AppLocalizations.of(context);

    if (value == null || value.isEmpty) {
      return translate?.empty_address ??
          'Le champ adresse ne peut pas être vide.';
    }
    return null;
  }

  void _presentDatePicker() async {
    final now = DateTime.now();
    final firstDate = DateTime(now.year, now.month, now.day);
    final lastDate = DateTime(now.year + 4);

    // Sélectionner la date
    final pickedDate = await showDatePicker(
      context: context,
      initialDate: now,
      firstDate: firstDate,
      lastDate: lastDate,
    );

    // Sélectionner l'heure
    final TimeOfDay? pickedTime = await showTimePicker(
      context: context,
      initialTime: TimeOfDay.now(),
    );

    if (pickedDate != null && pickedTime != null) {
      // Convertir l'heure en DateTime
      final pickedDateTime = DateTime(
        pickedDate.year,
        pickedDate.month,
        pickedDate.day,
        pickedTime.hour,
        pickedTime.minute,
      );

      // Formater la date et l'heure en ISO 8601
      final iso8601FormattedDateTime =
          DateFormat("yyyy-MM-ddTHH:mm:ss").format(pickedDateTime.toUtc()) + "Z";

      setState(() {
        _selectedDate =
            iso8601FormattedDateTime; // Assurez-vous d'utiliser le bon champ dans votre code
      });
    }
  }

  void _saveEvent() async {
    final translate = AppLocalizations.of(context);
    if (_formKey.currentState!.validate()) {
      _formKey.currentState!.save();
      if (_selectedDate == null || _selectedDate!.isEmpty) {
        ScaffoldMessenger.of(context).clearSnackBars();
        ScaffoldMessenger.of(context).showSnackBar(
          SnackBar(
            content: Text(
              textAlign: TextAlign.center,
              translate?.empty_date ?? "Veuillez renseigner une date.",
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
          "is_public": _isPublic,
          "latitude": _latitude,
          "longitude": _longitude,
        };
      await eventService.createEvent(newEvent);
        ScaffoldMessenger.of(context).clearSnackBars();
        context.go('/home');
        Future.delayed(Duration(milliseconds: 300), () {
          ScaffoldMessenger.of(context).showSnackBar(
            SnackBar(
              content: Text(
                textAlign: TextAlign.center,
                translate?.event_saved_success ??
                    "L'événement a bien été enregistré.",
                style:
                    TextStyle(color: Theme.of(context).colorScheme.onPrimary),
              ),
              backgroundColor: Theme.of(context).colorScheme.primary,
            ),
          );
        });
      } on AppException catch (error) {
        debugPrint(error.message);
        ScaffoldMessenger.of(context).clearSnackBars();
        ScaffoldMessenger.of(context).showSnackBar(
          SnackBar(
            content: Text(
              textAlign: TextAlign.center,
              translate?.server_error ??
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
      _addressController.dispose();
      super.dispose();
    }
  }

  @override
  Widget build(BuildContext context) {
    final translate = AppLocalizations.of(context);

    return Scaffold(
      appBar: AppBar(
        title: Text(translate?.create_event ?? "Créer un événement"),
      ),
      body: Center(
        child: SingleChildScrollView(
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
                            return translate?.fifty_char ??
                                'Doit contenir entre 1 et 50 caractères.';
                          }
                          return null;
                        },
                        style: TextStyle(
                            color: Theme.of(context).colorScheme.onSurface),
                        maxLength: 50,
                        decoration: InputDecoration(
                          border: const OutlineInputBorder(),
                          icon: const Icon(Icons.title),
                          label: Text(
                              translate?.event_name ?? 'Nom de l\'événement'),
                        ),
                        onSaved: (value) {
                          _enteredName = value!;
                        },
                        onTapOutside: (event) {
                          FocusScope.of(context).unfocus();
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
                            style: TextStyle(
                                color: Theme.of(context).colorScheme.onSurface),
                            decoration: InputDecoration(
                              border: OutlineInputBorder(),
                              icon: Icon(Icons.place),
                              labelText: translate?.event_address ??
                                  'Adresse de l\'événement',
                            ),
                            validator: _validateAddress,
                            onSaved: (value) {
                              _addressController.text = value!;
                            },
                            onTapOutside: (event) {
                              FocusScope.of(context).unfocus();
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
                          _addressController.text =
                              suggestion; // Mise à jour du champ de texte
                          final selectedFeature =
                              _suggestedAddresses.firstWhere((feature) =>
                                  feature['properties']['label'] == suggestion);
                          final latitude =
                              selectedFeature['geometry']['coordinates'][1];
                          final longitude =
                              selectedFeature['geometry']['coordinates'][0];

                          setState(() {
                            _latitude = latitude;
                            _longitude = longitude;
                          });
                        },
                      ),
                      SizedBox(height: 40),
                      Padding(
                        padding: EdgeInsets.symmetric(horizontal: 2),
                        child: Row(
                          children: [
                            Icon(Icons.public),
                            SizedBox(width: 15),
                            Checkbox(
                              value: _isPublic,
                              // Utilise la valeur actuelle de _isPublic
                              onChanged: (bool? value) {
                                setState(() {
                                  _isPublic = value!;
                                });
                              },
                            ),
                            Text(translate?.open_public ?? "Ouvert au public"),
                            SizedBox(width: 20),
                          ],
                        ),
                      ),
                      SizedBox(height: 15),
                      Padding(
                        padding: EdgeInsets.symmetric(horizontal: 2),
                        child: Row(
                          children: [
                            Icon(Icons.calendar_month),
                            SizedBox(width: 15),
                            ElevatedButton(
                              onPressed: _presentDatePicker,
                              child: Text(translate?.select_date ??
                                  "Sélectionner une date"),
                            ),
                            SizedBox(width: 20),
                            Text(
                              style: TextStyle(
                                  color:
                                      Theme.of(context).colorScheme.onSurface),
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
                                    return translate?.empty_type ??
                                        'Veuillez sélectionner un type.';
                                  }
                                  return null;
                                },
                                hint: Text(
                                  translate?.type_select_label ?? "Type",
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
                                      translate?.type_select_label ?? "Type",
                                      style: TextStyle(
                                          color: Theme.of(context)
                                              .colorScheme
                                              .onSurface),
                                    ),
                                  ),
                                  DropdownMenuItem(
                                    value: "match",
                                    child: Text(
                                      translate?.match ?? "Match",
                                      style: TextStyle(
                                          color: Theme.of(context)
                                              .colorScheme
                                              .onSurface),
                                    ),
                                  ),
                                  DropdownMenuItem(
                                    value: "training",
                                    child: Text(
                                      translate?.training ?? "Training",
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
                                    return translate?.empty_sport ??
                                        'Veuillez sélectionner un sport.';
                                  }
                                  return null;
                                },
                                hint: Text(
                                  translate?.sport_select_label ?? "Sport",
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
                                      translate?.sport_select_label ?? "Sport",
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
                      SizedBox(
                        height: 60,
                      ),
                      ElevatedButton(
                          onPressed: _saveEvent,
                          child: Text(translate?.save_event ?? "Enregistrer"))
                    ],
                  ),
                ),
              )
            ],
          ),
        ),
      ),
    );
  }
}
