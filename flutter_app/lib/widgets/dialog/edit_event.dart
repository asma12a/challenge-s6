import 'package:flutter/material.dart';
import 'package:squad_go/core/models/event.dart';
import 'package:squad_go/core/services/event_service.dart';

class EditEventDialog extends StatefulWidget {
  final Event event;
  final Future<void> Function()? onRefresh;

  const EditEventDialog({super.key, required this.event, this.onRefresh});

  @override
  State<EditEventDialog> createState() => _EditEventDialogState();
}

class _EditEventDialogState extends State<EditEventDialog> {
  final EventService eventService = EventService();

  @override
  Widget build(BuildContext context) {
    return const Placeholder();
  }
}
