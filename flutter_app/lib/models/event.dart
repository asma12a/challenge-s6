class Event {
  const Event({
    this.id,
    this.eventType,
    required this.name,
    required this.address,
    required this.date,
    required this.sport,
  });

  final String? id;
  final String name;
  final String address;
  final String date;
  final String sport;
  final String? eventType;

  Map<String, dynamic> toJson() {
    return {
      'name': name,
      'address': address,
      'date': date,
      'sport_id': sport,
      'event_type': eventType,
    };
  }
}