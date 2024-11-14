import 'package:flutter_app/models/sport.dart';

class Event {
  const Event(
      {required this.id,
      required this.name,
      required this.address,
      required this.date,
      required this.sport});

  final String id;
  final String name;
  final String address;
  final String date;
  final Sport sport;
}
