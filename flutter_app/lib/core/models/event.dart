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
    this.latitude = 46.603354,
    this.longitude = 1.888334,
    this.isPublic = true,
    this.createdBy,
    this.teams,
    this.type = EventType.match,
  });

  final String? id;
  final String? createdBy;
  final String name;
  final String address;
  final double latitude;
  final double longitude;
  final String date;
  final Sport sport;
  final EventType type;
  final bool isPublic;
  final List<Team>? teams;

  factory Event.fromJson(Map<String, dynamic> data) {
    return Event(
      id: data['id'],
      name: data['name'],
      address: data['address'],
      latitude: data['latitude'] ?? 46.603354,
      longitude: data['longitude'] ?? 1.888334,
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
      isPublic: data['is_public'] ?? true,
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'id': id,
      'name': name,
      'address': address,
      'latitude': latitude,
      'longitude': longitude,
      'date': DateFormat('yyyy-MM-dd').format(DateTime.parse(date)),
      'sport_id': sport.id,
      'event_type': type == EventType.match ? 'match' : 'training',
      'is_public': isPublic,
    };
  }

  Event copyWith({
    String? id,
    String? name,
    String? address,
    double? latitude,
    double? longitude,
    String? date,
    Sport? sport,
    EventType? type,
    List<Team>? teams,
    String? createdBy,
    bool? isPublic,
  }) {
    return Event(
      id: id ?? this.id,
      name: name ?? this.name,
      address: address ?? this.address,
      latitude: latitude ?? this.latitude,
      longitude: longitude ?? this.longitude,
      date: date ?? this.date,
      sport: sport ?? this.sport,
      type: type ?? this.type,
      teams: teams ?? this.teams,
      createdBy: createdBy ?? this.createdBy,
      isPublic: isPublic ?? this.isPublic,
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
