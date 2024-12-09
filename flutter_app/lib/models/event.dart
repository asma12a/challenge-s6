import 'package:flutter/material.dart';
import 'package:squad_go/models/sport.dart';
import 'package:squad_go/models/team.dart';

class Event {
  const Event({
    required this.id,
    required this.name,
    required this.address,
    required this.date,
    required this.sport,
    this.teams,
    this.type = EventType.match,
  });

  final String id;
  final String name;
  final String address;
  final String date;
  final Sport sport;
  final EventType type;
  final List<Team>? teams;
}

enum EventType { match, training }

final Map<EventType, Color> eventTypeColor = {
  EventType.match: Colors.red,
  EventType.training: Colors.blue,
};

final Map<EventType, IconData> eventTypeIcon = {
  EventType.match: Icons.sports,
  EventType.training: Icons.fitness_center,
};
