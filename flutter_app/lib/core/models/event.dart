import 'package:flutter/material.dart';
import 'package:squad_go/core/models/sport.dart';
import 'package:squad_go/core/models/team.dart';

class Event {
  const Event({
    this.id,
    required this.name,
    required this.address,
    required this.date,
    required this.sport,
    this.teams,
    this.type = EventType.match,
  });

  final String? id;
  final String name;
  final String address;
  final String date;
  final Sport sport;
  final EventType type;
  final List<Team>? teams;

  factory Event.fromJson(Map<String, dynamic> data) {
    return Event(
      id: data['id'],
      name: data['name'],
      address: data['address'],
      date: data['date'],
      sport: Sport.fromJson(data['sport']),
      type:
          data['event_type'] == 'match' ? EventType.match : EventType.training,
      teams: data['teams'] != null
          ? (data['teams'] as List<dynamic>)
              .map((team) => Team.fromJson(team))
              .toList()
          : null,
    );
  }
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
