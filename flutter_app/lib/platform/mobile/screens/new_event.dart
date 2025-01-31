import 'package:animated_toggle_switch/animated_toggle_switch.dart';
import 'package:squad_go/core/exceptions/app_exception.dart';
import 'package:squad_go/core/models/event.dart';
import 'package:squad_go/core/models/sport.dart';
import 'package:squad_go/core/services/event_service.dart';
import 'package:flutter_typeahead/flutter_typeahead.dart';
import 'package:flutter/material.dart';
import 'package:intl/intl.dart';
import 'package:squad_go/main.dart';
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
  String iso8601FormattedDateTime = "";
  final TextEditingController _addressController = TextEditingController();
  final TextEditingController _dateController = TextEditingController();
  List<dynamic> _suggestedAddresses = [];
  EventType? _selectedType;
  String? _selectedSport;
  String? showedDate;
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

      setState(() {
        iso8601FormattedDateTime =
            '${DateFormat("yyyy-MM-ddTHH:mm:ss").format(pickedDateTime.toUtc())}Z';
        _selectedDate = iso8601FormattedDateTime;
        _dateController.text =
            DateFormat("dd/MM/yyyy HH:mm").format(pickedDateTime);
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
          "event_type": _selectedType?.name,
          "is_public": _isPublic,
          "latitude": _latitude,
          "longitude": _longitude,
        };
        await eventService.createEvent(newEvent);
        ScaffoldMessenger.of(context).clearSnackBars();

        // Afficher le Snackbar
        ScaffoldMessenger.of(context).showSnackBar(
          SnackBar(
            content: Text(
              textAlign: TextAlign.center,
              translate?.event_saved_success ??
                  "L'événement a bien été enregistré.",
              style: TextStyle(color: Theme.of(context).colorScheme.onPrimary),
            ),
            backgroundColor: Theme.of(context).colorScheme.primary,
          ),
        );
        context.go('/home');
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
                          prefixIcon: const Icon(Icons.title),
                          labelText:
                              translate?.event_name ?? 'Nom de l\'événement',
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
                              prefixIcon: Icon(Icons.place),
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
                      SizedBox(height: 30),
                      TextFormField(
                        controller: _dateController,
                        readOnly: true,
                        onTap: () async {
                          _presentDatePicker();
                        },
                        onTapOutside: (event) {
                          FocusScope.of(context).unfocus();
                        },
                        validator: (value) {
                          if (value == null || value.isEmpty) {
                            return translate?.empty_date ??
                                'Veuillez sélectionner une date';
                          }
                          return null;
                        },
                        decoration: InputDecoration(
                            border: OutlineInputBorder(),
                            prefixIcon: Icon(Icons.calendar_month),
                            labelText: translate?.event_date ??
                                'Date de l\'événement'),
                      ),
                      SizedBox(height: 30),
                      Row(
                        children: [
                          Expanded(
                            child: DropdownButtonFormField<String>(
                              value: _selectedSport,
                              decoration: InputDecoration(
                                border: OutlineInputBorder(),
                                labelText:
                                    translate?.sport_select_label ?? 'Sport',
                              ),
                              items: _sports.map((sport) {
                                return DropdownMenuItem(
                                  value: sport.id,
                                  child: Row(
                                    children: [
                                      if (sportIcon.containsKey(sport.name))
                                        Padding(
                                          padding:
                                              const EdgeInsets.only(right: 8.0),
                                          child: Icon(
                                            sportIcon[sport.name],
                                            size: 16,
                                          ),
                                        ),
                                      Text(sport.name.name[0].toUpperCase() +
                                          sport.name.name.substring(1)),
                                    ],
                                  ),
                                );
                              }).toList(),
                              onChanged: (value) {
                                setState(() {
                                  _selectedSport = value!;
                                });
                              },
                              validator: (value) {
                                if (value == null) {
                                  return translate?.empty_sport ??
                                      'Veuillez sélectionner un sport';
                                }
                                return null;
                              },
                              onSaved: (value) {
                                _selectedSport = value!;
                              },
                            ),
                          ),
                          SizedBox(width: 16),
                          Expanded(
                            child: DropdownButtonFormField<EventType>(
                              value: _selectedType,
                              decoration: InputDecoration(
                                  border: OutlineInputBorder(),
                                  labelText: translate?.event_type ??
                                      'Type d\'événement'),
                              items: EventType.values
                                  .map((type) => DropdownMenuItem(
                                        value: type,
                                        child: Row(
                                          children: [
                                            if (eventTypeIcon.containsKey(type))
                                              Padding(
                                                padding: const EdgeInsets.only(
                                                    right: 8.0),
                                                child: Icon(
                                                  eventTypeIcon[type],
                                                  size: 16,
                                                ),
                                              ),
                                            Text(type.name[0].toUpperCase() +
                                                type.name.substring(1)),
                                          ],
                                        ),
                                      ))
                                  .toList(),
                              onChanged: (value) {
                                setState(() {
                                  _selectedType = value!;
                                });
                              },
                              onSaved: (value) {
                                if (value != null) {
                                  _selectedType = value;
                                }
                              },
                            ),
                          ),
                        ],
                      ),
                      SizedBox(height: 30),
                      AnimatedToggleSwitch.dual(
                        first: false,
                        second: true,
                        current: _isPublic,
                        onChanged: (value) => setState(() {
                          _isPublic = value;
                        }),
                        iconBuilder: (value) => value
                            ? const Icon(
                                Icons.public,
                                color: Colors.white,
                              )
                            : const Icon(
                                Icons.lock,
                                color: Colors.white,
                              ),
                        height: 40,
                        style: ToggleStyle(
                          indicatorColor: Colors.blue,
                          borderColor: Colors.blue,
                        ),
                        textBuilder: (value) => value
                            ? Text(
                                translate?.public ?? 'Public',
                                style: TextStyle(fontWeight: FontWeight.bold),
                              )
                            : Text(
                                translate?.private ?? 'Privé',
                                style: TextStyle(fontWeight: FontWeight.bold),
                              ),
                      ),
                      SizedBox(
                        height: 60,
                      ),
                      ElevatedButton.icon(
                        onPressed: _saveEvent,
                        icon: Icon(
                          Icons.save,
                          color: Colors.white,
                        ),
                        label: Text(translate?.save_event ?? "Enregistrer"),
                        style: ElevatedButton.styleFrom(
                          foregroundColor: Colors.white,
                          backgroundColor: Colors.blue,
                          padding: const EdgeInsets.symmetric(
                              horizontal: 24, vertical: 12),
                          textStyle: const TextStyle(
                            fontSize: 18,
                            fontWeight: FontWeight.bold,
                          ),
                        ),
                      )
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
