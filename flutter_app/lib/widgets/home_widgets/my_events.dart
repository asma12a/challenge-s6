import 'package:flutter/material.dart';
import 'package:squad_go/core/models/event.dart';
import 'package:squad_go/core/services/event_service.dart';
import 'package:squad_go/main.dart';
import 'package:squad_go/widgets/carousel.dart';
import 'package:squad_go/widgets/event_card.dart';

class HomeMyEvents extends StatefulWidget {
  const HomeMyEvents({super.key});

  @override
  State<HomeMyEvents> createState() => HomeMyEventsState();
}

class HomeMyEventsState extends State<HomeMyEvents> {
  final EventService eventService = EventService();
  List<Event> myEvents = [];

  @override
  void initState() {
    super.initState();

    fetchMyEvents();
  }

  Future<void> fetchMyEvents() async {
    try {
      List<Event> events = await eventService.getMyEvents();
      DateTime now = DateTime.now();
      final filteredEvents = events.where((event) {
        DateTime eventDate = DateTime.parse(event.date);
        return eventDate.isAfter(now.subtract(Duration(days: 1)));
      }).toList();

      setState(() {
        myEvents = filteredEvents;
      });
    } catch (e) {
      // Handle error
      log.severe('Failed to fetch events: $e');
    }
  }

  @override
  Widget build(BuildContext context) {
    return Carousel(
      text: "Mes événements en cours",
      items: myEvents
          .map((event) => EventCard(
                event: event,
                hasJoinedEvent: true,
              ))
          .toList(),
    );
  }
}
