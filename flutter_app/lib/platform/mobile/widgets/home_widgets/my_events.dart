import 'package:flutter/material.dart';
import 'package:squad_go/core/models/event.dart';
import 'package:squad_go/core/services/event_service.dart';
import 'package:squad_go/main.dart';
import 'package:squad_go/platform/mobile/widgets/carousel.dart';
import 'package:squad_go/platform/mobile/widgets/event_card.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';

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
    final translate = AppLocalizations.of(context);
    return Carousel(
      text: translate?.my_events ?? "Mes événements en cours",
      items: myEvents
          .map((event) => EventCard(
                event: event,
                hasJoinedEvent: true,
              ))
          .toList(),
    );
  }
}