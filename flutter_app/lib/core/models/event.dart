import 'package:flutter/material.dart';
import 'package:intl/intl.dart';
import 'package:squad_go/core/models/sport.dart';
import 'package:squad_go/core/models/team.dart';

class Event {
  const Event({
    this.id,
    required this.name,
    required this.address,
    required this.date,
    required this.sport,
    this.createdBy,
    this.teams,
    this.type = EventType.match,
  });

  final String? id;
  final String? createdBy;
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
      createdBy: data['created_by'],
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'id': id,
      'name': name,
      'address': address,
      'date': DateFormat('yyyy-MM-dd').format(DateTime.parse(date)),
      'sport_id': sport.id,
      'event_type': type == EventType.match ? 'match' : 'training',
    };
  }

  Event copyWith({
    String? id,
    String? name,
    String? address,
    String? date,
    Sport? sport,
    EventType? type,
    List<Team>? teams,
    String? createdBy,
  }) {
    return Event(
      id: id ?? this.id,
      name: name ?? this.name,
      address: address ?? this.address,
      date: date ?? this.date,
      sport: sport ?? this.sport,
      type: type ?? this.type,
      teams: teams ?? this.teams,
      createdBy: createdBy ?? this.createdBy,
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
