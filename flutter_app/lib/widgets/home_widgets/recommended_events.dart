import 'package:flutter/material.dart';
import 'package:squad_go/core/models/event.dart';
import 'package:squad_go/core/services/event_service.dart';
import 'package:squad_go/main.dart';
import 'package:squad_go/widgets/carousel.dart';
import 'package:squad_go/widgets/event_card.dart';

class HomeRecommendedEvents extends StatefulWidget {
  const HomeRecommendedEvents({super.key});

  @override
  State<HomeRecommendedEvents> createState() => HomeRecommendedEventsState();
}

class HomeRecommendedEventsState extends State<HomeRecommendedEvents> {
  final EventService eventService = EventService();
  List<Event> recommendedEvents = [];

  @override
  void initState() {
    super.initState();

    fetchRecommendedEvents();
  }

  Future<void> fetchRecommendedEvents() async {
    try {
      List<Event> events = await eventService.getRecommendedEvents();

      setState(() {
        recommendedEvents = events;
      });
    } catch (e) {
      // Handle error
      log.severe('Failed to fetch events: $e');
    }
  }

  @override
  Widget build(BuildContext context) {
    return Carousel(
      text: "Évenements recommandés",
      items: recommendedEvents
          .map((event) => EventCard(
                event: event,
                hasJoinedEvent: false,
              ))
          .toList(),
    );
  }
}
