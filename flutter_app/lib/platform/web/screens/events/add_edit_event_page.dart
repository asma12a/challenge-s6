import 'dart:ui';
import 'package:flutter/material.dart';
import 'package:squad_go/core/services/event_service.dart';

class AddEditEventModal extends StatefulWidget {
  final Map<String, dynamic>? event;
  final VoidCallback? onEventSaved;

  const AddEditEventModal({super.key, this.event, this.onEventSaved});

  @override
  State<AddEditEventModal> createState() => _AddEditEventModalState();
}

class _AddEditEventModalState extends State<AddEditEventModal> {
  final EventService _eventService = EventService();
  final _formKey = GlobalKey<FormState>();
  late String name;
  late String date;
  late String eventCode;
  late bool isPublic;
  late bool isFinished;
  late String address;

  @override
  void initState() {
    super.initState();
    name = widget.event?['name'] ?? '';
    date = widget.event?['date'] ?? '';
    eventCode = widget.event?['event_code'] ?? '';
    isPublic = widget.event?['is_public'] ?? false;
    isFinished = widget.event?['is_finished'] ?? false;
    address = widget.event?['address'] ?? '';
  }

  Future<void> saveEvent() async {
    if (!_formKey.currentState!.validate()) return;

    _formKey.currentState!.save();
    try {
      if (widget.event == null) {
        await _eventService.createEvent({
          'name': name,
          'date': date,
          'event_code': eventCode,
          'is_public': isPublic.toString(),
          'is_finished': isFinished.toString(),
          'address': address,
        });
      } else {
        await _eventService.updateEvent(
          widget.event!['id'],
          {
            'name': name,
            'date': date,
            'event_code': eventCode,
            'is_public': isPublic,
            'is_finished': isFinished,
            'address': address,
          },
        );
      }

      if (widget.onEventSaved != null) widget.onEventSaved!();
      Navigator.pop(context);
    } catch (e) {
      ScaffoldMessenger.of(context).showSnackBar(
        SnackBar(content: Text('Erreur: $e')),
      );
    }
  }

  @override
  Widget build(BuildContext context) {
    return BackdropFilter(
      filter: ImageFilter.blur(sigmaX: 10, sigmaY: 10),
      child: Dialog(
        shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(16)),
        backgroundColor: Colors.white,
        child: Container(
          padding: const EdgeInsets.all(16),
          width: MediaQuery.of(context).size.width * 0.8,
          constraints: const BoxConstraints(maxWidth: 400),
          child: Form(
            key: _formKey,
            child: Column(
              mainAxisSize: MainAxisSize.min,
              crossAxisAlignment: CrossAxisAlignment.stretch,
              children: [
                Align(
                  alignment: Alignment.topRight,
                  child: IconButton(
                    icon: const Icon(Icons.close),
                    onPressed: () => Navigator.pop(context),
                  ),
                ),
                Text(
                  widget.event == null
                      ? 'Ajouter un Événement'
                      : 'Modifier l\'Événement',
                  textAlign: TextAlign.center,
                  style: const TextStyle(
                    fontSize: 18,
                    fontWeight: FontWeight.bold,
                  ),
                ),
                const SizedBox(height: 20),
                TextFormField(
                  initialValue: name,
                  decoration: InputDecoration(
                    labelText: 'Nom',
                    border: OutlineInputBorder(
                      borderRadius: BorderRadius.circular(8),
                    ),
                  ),
                  validator: (value) =>
                      value!.isEmpty ? 'Veuillez entrer un nom.' : null,
                  onSaved: (value) => name = value!,
                ),
                const SizedBox(height: 16),
                TextFormField(
                  initialValue: date,
                  decoration: InputDecoration(
                    labelText: 'Date',
                    border: OutlineInputBorder(
                      borderRadius: BorderRadius.circular(8),
                    ),
                  ),
                  validator: (value) =>
                      value!.isEmpty ? 'Veuillez entrer une date.' : null,
                  onSaved: (value) => date = value!,
                ),
                const SizedBox(height: 16),
                TextFormField(
                  initialValue: eventCode,
                  decoration: InputDecoration(
                    labelText: 'Code de l\'événement',
                    border: OutlineInputBorder(
                      borderRadius: BorderRadius.circular(8),
                    ),
                  ),
                  validator: (value) =>
                      value!.isEmpty ? 'Veuillez entrer un code.' : null,
                  onSaved: (value) => eventCode = value!,
                ),
                const SizedBox(height: 16),
                TextFormField(
                  initialValue: address,
                  decoration: InputDecoration(
                    labelText: 'Adresse',
                    border: OutlineInputBorder(
                      borderRadius: BorderRadius.circular(8),
                    ),
                  ),
                  onSaved: (value) => address = value!,
                ),
                const SizedBox(height: 16),
                Row(
                  mainAxisAlignment: MainAxisAlignment.spaceBetween,
                  children: [
                    TextButton(
                      onPressed: () => Navigator.pop(context),
                      style: TextButton.styleFrom(
                        foregroundColor: Colors.black,
                        backgroundColor: Colors.white,
                        side: const BorderSide(color: Colors.black),
                        padding: const EdgeInsets.symmetric(
                            horizontal: 24, vertical: 12),
                        shape: RoundedRectangleBorder(
                          borderRadius: BorderRadius.circular(8),
                        ),
                      ),
                      child: const Text('Annuler'),
                    ),
                    ElevatedButton(
                      onPressed: saveEvent,
                      style: ElevatedButton.styleFrom(
                        foregroundColor:
                            const Color.fromARGB(255, 255, 255, 255),
                        backgroundColor: Theme.of(context).primaryColor,
                        padding: const EdgeInsets.symmetric(
                            horizontal: 24, vertical: 12),
                        shape: RoundedRectangleBorder(
                          borderRadius: BorderRadius.circular(8),
                        ),
                      ),
                      child: const Text('Enregistrer'),
                    ),
                  ],
                ),
              ],
            ),
          ),
        ),
      ),
    );
  }
}
