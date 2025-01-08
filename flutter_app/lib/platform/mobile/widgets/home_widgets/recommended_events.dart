import 'package:flutter/material.dart';
import 'package:squad_go/core/models/event.dart';
import 'package:squad_go/core/services/event_service.dart';
import 'package:squad_go/main.dart';
import 'package:squad_go/platform/mobile/widgets/carousel.dart';
import 'package:squad_go/platform/mobile/widgets/event_card.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';

class HomeRecommendedEvents extends StatefulWidget {
  final Future<void> Function()? onRefresh;
  const HomeRecommendedEvents({super.key, this.onRefresh});

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
    final translate = AppLocalizations.of(context);
    return Carousel(
      text: translate?.recommended_events ?? "Évenements recommandés",
      items: recommendedEvents
          .map((event) => EventCard(
                event: event,
                hasJoinedEvent: false,
                onRefresh: widget.onRefresh,
              ))
          .toList(),
    );
  }
}
