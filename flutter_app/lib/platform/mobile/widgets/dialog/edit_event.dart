import 'dart:convert';

import 'dart:convert';

import 'package:animated_toggle_switch/animated_toggle_switch.dart';
import 'package:flutter/material.dart';
import 'package:flutter_typeahead/flutter_typeahead.dart';
import 'package:intl/intl.dart';
import 'package:squad_go/core/models/event.dart';
import 'package:squad_go/core/models/sport.dart';
import 'package:squad_go/core/services/event_service.dart';
import 'package:squad_go/main.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';

class EditEventDialog extends StatefulWidget {
  final Event event;
  final Future<void> Function()? onRefresh;

  const EditEventDialog({super.key, required this.event, this.onRefresh});

  @override
  State<EditEventDialog> createState() => _EditEventDialogState();
}

class _EditEventDialogState extends State<EditEventDialog> {
  final EventService eventService = EventService();
  late Event event;

  final _formKey = GlobalKey<FormState>();
  final TextEditingController _addressController = TextEditingController();
  final TextEditingController _dateController = TextEditingController();
  String iso8601FormattedDateTime = "";
  List<dynamic> _suggestedAddresses = [];
  List<Sport> _sports = [];

  @override
  void initState() {
    super.initState();
    event = widget.event;
    iso8601FormattedDateTime = event.date;
    debugPrint("Date ISO : $iso8601FormattedDateTime");
    _addressController.text = event.address;
    _dateController.text = DateFormat('yyyy/MM/dd HH:mm').format(
      DateTime.parse(event.date).add(
        Duration(hours: 1),
      ),
    );
    _fetchSports();
  }

  String? _validateAddress(String? value) {
    final translate = AppLocalizations.of(context);
    if (value == null || value.isEmpty) {
      return translate?.empty_address ??
          'Le champ adresse ne peut pas être vide.';
    }
    return null;
  }

  Future<void> _fetchSports() async {
    try {
      final sports = await eventService.getSports();

      setState(() {
        _sports = sports;
      });
    } catch (e) {
      // Handle error
      log.severe('Failed to fetch sports: $e');
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

  void _presentDatePicker() async {
    final now = DateTime.now();
    final firstDate = DateTime(now.year, now.month, now.day);
    final lastDate = DateTime(now.year + 4);

    final initialDate = event.date.isNotEmpty
        ? DateFormat('yyyy-MM-dd').parse(event.date)
        : now;

    final pickedDate = await showDatePicker(
      context: context,
      initialDate: initialDate,
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
            "${DateFormat("yyyy-MM-ddTHH:mm:ss").format(pickedDateTime.toUtc())}Z";
        // Met à jour le champ de texte avec la date et l'heure formatée
        _dateController.text =
            DateFormat('yyyy/MM/dd HH:mm').format(pickedDateTime);
      });
    }
  }

  void _updateEvent() async {
    final translate = AppLocalizations.of(context);
    if (_formKey.currentState!.validate()) {
      _formKey.currentState!.save();
      try {
        debugPrint("Event to update : ${event.toJson()}");
        await eventService.updateEvent(event.id!, event.toJson());
        ScaffoldMessenger.of(context).clearSnackBars();
        Navigator.of(context).pop();
        ScaffoldMessenger.of(context).showSnackBar(
          SnackBar(
            content: Text(
              translate?.event_updated_success ??
                  "L'événement a bien été mis à jour.",
              textAlign: TextAlign.center,
              style: TextStyle(
                color: Theme.of(context).colorScheme.onPrimary,
              ),
            ),
            backgroundColor: Theme.of(context).colorScheme.primary,
          ),
        );
        widget.onRefresh
            ?.call(); // Cette ligne doit être en dehors de `Future.delayed`
      } catch (e) {
        // Gestion des erreurs
        // Gestion des erreurs
        log.severe('Failed to update event: $e');
      }
    }
  }

  @override
  Widget build(BuildContext context) {
    final translate = AppLocalizations.of(context);
    return Dialog(
      child: Container(
        padding: const EdgeInsets.all(16),
        child: Form(
          key: _formKey,
          child: SingleChildScrollView(
            child: Column(
              mainAxisSize: MainAxisSize.min,
              children: [
                Text(
                  '${translate?.edit ?? "Modifier:"} ${event.name ?? ""}',
                  maxLines: 1,
                  overflow: TextOverflow.ellipsis,
                  style: TextStyle(
                    fontSize: 20,
                    fontWeight: FontWeight.bold,
                  ),
                ),
                SizedBox(height: 8),
                AnimatedToggleSwitch.dual(
                  first: false,
                  second: true,
                  current: event.isPublic,
                  onChanged: (value) => setState(() {
                    event = event.copyWith(isPublic: value);
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
                SizedBox(height: 8),
                TextFormField(
                  initialValue: event.name,
                  onTapOutside: (event) {
                    FocusScope.of(context).unfocus();
                  },
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
                  maxLength: 50,
                  decoration: InputDecoration(
                      labelText:
                          translate?.event_name ?? 'Nom de l\'événement'),
                  onSaved: (value) {
                    event = event.copyWith(name: value);
                  },
                ),
                TypeAheadField<String>(
                  controller: _addressController,
                  builder: (context, controller, focusNode) {
                    return TextFormField(
                      controller: controller,
                      focusNode: focusNode,
                      decoration: InputDecoration(
                          labelText: translate?.event_address ??
                              'Adresse de l\'événement'),
                      validator: _validateAddress,
                      onSaved: (value) {
                        _addressController.text = value!;
                        event = event.copyWith(address: value);
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
                    _addressController.text = suggestion;
                    final selectedFeature = _suggestedAddresses.firstWhere(
                        (feature) =>
                            feature['properties']['label'] == suggestion);
                    final latitude =
                        selectedFeature['geometry']['coordinates'][1];
                    final longitude =
                        selectedFeature['geometry']['coordinates'][0];
                    event = event.copyWith(
                      address: suggestion,
                      latitude: latitude,
                      longitude: longitude,
                    );
                  },
                ),
                SizedBox(height: 16),
                TextFormField(
                  controller: _dateController,
                  readOnly: true,
                  onTap: () async {
                    _presentDatePicker();
                    setState(() {
                      _dateController.text = DateFormat('dd/MM/yyyy')
                          .format(DateTime.parse(event.date));
                    });
                  },
                  validator: (value) {
                    if (value == null || value.isEmpty) {
                      return translate?.empty_date ??
                          'Veuillez sélectionner une date';
                    }
                    return null;
                  },
                  decoration: InputDecoration(
                      labelText:
                          translate?.event_date ?? 'Date de l\'événement'),
                  onSaved: (value) {
                    event = event.copyWith(date: iso8601FormattedDateTime);
                  },
                ),
                Row(
                  children: [
                    Expanded(
                      child: DropdownButtonFormField<String>(
                        value: event.sport.id,
                        decoration: InputDecoration(
                            labelText:
                                translate?.sport_select_label ?? 'Sport'),
                        items: _sports
                            .map((sport) => DropdownMenuItem(
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
                                ))
                            .toList(),
                        onChanged: (value) {
                          setState(() {
                            event = event.copyWith(
                                sport: _sports
                                    .firstWhere((sport) => sport.id == value));
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
                          event = event.copyWith(
                              sport: _sports
                                  .firstWhere((sport) => sport.id == value));
                        },
                      ),
                    ),
                    SizedBox(width: 16),
                    Expanded(
                      child: DropdownButtonFormField<EventType>(
                        value: event.type,
                        decoration: InputDecoration(
                            labelText:
                                translate?.event_type ?? 'Type d\'événement'),
                        items: EventType.values
                            .map((type) => DropdownMenuItem(
                                  value: type,
                                  child: Row(
                                    children: [
                                      if (eventTypeIcon.containsKey(type))
                                        Padding(
                                          padding:
                                              const EdgeInsets.only(right: 8.0),
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
                            event = event.copyWith(type: value);
                          });
                        },
                        validator: (value) {
                          if (value == null) {
                            return translate?.empty_type ??
                                'Veuillez sélectionner un type de match';
                          }
                          return null;
                        },
                        onSaved: (value) {
                          event = event.copyWith(type: value);
                        },
                      ),
                    ),
                  ],
                ),
                SizedBox(height: 16),
                ElevatedButton.icon(
                  onPressed: () {
                    _updateEvent();
                  },
                  icon: const Icon(
                    Icons.update,
                    color: Colors.white,
                  ),
                  label: Text(translate?.update ?? 'Mettre à jour'),
                  style: ElevatedButton.styleFrom(
                    foregroundColor: Colors.white,
                    backgroundColor: Colors.blue,
                    padding: const EdgeInsets.symmetric(
                        horizontal: 20, vertical: 10),
                    textStyle: const TextStyle(
                      fontSize: 16,
                      fontWeight: FontWeight.bold,
                    ),
                  ),
                ),
              ],
            ),
          ),
        ),
      ),
    );
  }
}
