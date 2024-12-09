import 'package:squad_go/models/sport.dart';

class Event {
  const Event({
    this.id,
    required this.name,
    required this.address,
    required this.date,
    required this.sport,
    this.type = EventType.match,
  });

  final String? id;
  final String name;
  final String address;
  final String date;
  final Sport sport;
  final EventType type;
}

enum EventType { match, training }
