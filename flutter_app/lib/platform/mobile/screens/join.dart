import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:squad_go/core/models/event.dart';
import 'package:squad_go/core/services/event_service.dart';
import 'package:squad_go/platform/mobile/widgets/event_card.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';

class UpperCaseTextFormatter extends TextInputFormatter {
  @override
  TextEditingValue formatEditUpdate(
      TextEditingValue oldValue, TextEditingValue newValue) {
    return TextEditingValue(
      text: newValue.text.toUpperCase(),
      selection: newValue.selection,
    );
  }
}

class JoinEventScreen extends StatefulWidget {
  const JoinEventScreen({super.key});

  @override
  State<JoinEventScreen> createState() => _JoinEventScreenState();
}

class _JoinEventScreenState extends State<JoinEventScreen> {
  final EventService _eventService = EventService();
  final _formKey = GlobalKey<FormState>();
  var _enteredCode = '';
  Event? _event;
  bool hasSearched = false;

  void _joinEvent() async {
    final translate = AppLocalizations.of(context);
    if (_formKey.currentState!.validate()) {
      _formKey.currentState!.save();
      try {
        var event = await _eventService.getEventByCode(_enteredCode);

        setState(() {
          _event = event;
        });
      } catch (error) {
        setState(() {
          _event = null;
        });
        ScaffoldMessenger.of(context).clearSnackBars();
        ScaffoldMessenger.of(context).showSnackBar(
          SnackBar(
            content: Text(
              textAlign: TextAlign.center,
              translate?.no_event_code ??
                  "Aucun événement ne correspond à ce code.",
              style: TextStyle(color: Theme.of(context).colorScheme.onPrimary),
            ),
            backgroundColor: Theme.of(context).colorScheme.primary,
          ),
        );
      }
      setState(() {
        hasSearched = true;
      });
    }
  }

  @override
  Widget build(BuildContext context) {
    final translate = AppLocalizations.of(context);
    return Padding(
      padding: EdgeInsets.all(20),
      child: Column(
        children: [
          Text(
            translate?.join_event ?? "Rejoindre un événement",
            style: TextStyle(fontSize: 22),
          ),
          SizedBox(height: 45),
          Form(
            key: _formKey,
            child: TextFormField(
              validator: (value) {
                if (value == null || value.isEmpty) {
                  return translate?.empty_val ?? 'Veuillez saisir une valeur.';
                }
                if (value.length < 6) {
                  return translate?.six_char_code ?? 'Le code doit contenir 6 caractères.';
                }
                return null;
              },
              decoration: InputDecoration(
                border: const OutlineInputBorder(),
                icon: const Icon(Icons.qr_code),
                label: Text(
                    translate?.input_code ?? 'Saisir le code de l\'événement'),
              ),
              onSaved: (value) {
                _enteredCode = value!.toUpperCase();
              },
              textCapitalization: TextCapitalization.characters,
              maxLength: 6,
              onTapOutside: (event) => FocusScope.of(context).unfocus(),
              inputFormatters: [
                UpperCaseTextFormatter(),
              ],
            ),
          ),
          SizedBox(height: 30),
          ElevatedButton(
            onPressed: _formKey.currentState == null ||
                    !_formKey.currentState!.validate()
                ? null
                : _joinEvent,
            style: ElevatedButton.styleFrom(
              foregroundColor: Colors.white,
              backgroundColor: Colors.blue,
              padding: const EdgeInsets.symmetric(horizontal: 20, vertical: 10),
              textStyle: const TextStyle(
                fontSize: 16,
                fontWeight: FontWeight.bold,
              ),
            ),
            child: Text(translate?.search_code ?? "Rechercher par code"),
          ),
          SizedBox(height: 30),
          if (hasSearched && _event != null)
            EventCard(event: _event!, hasJoinedEvent: _event!.hasJoined),
          if (hasSearched && _event == null)
            Text(translate?.no_events ??
                "Aucun événement ne correspond à ce code."),
        ],
      ),
    );
  }
}
