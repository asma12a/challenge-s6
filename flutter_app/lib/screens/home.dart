import 'package:flutter/material.dart';
import 'package:squad_go/core/models/event.dart';
import 'package:squad_go/main.dart';
import 'package:squad_go/widgets/carousel.dart';
import 'package:squad_go/widgets/event_card.dart';
import 'package:squad_go/core/services/event_service.dart'; // Import EventService

class HomeScreen extends StatefulWidget {
  const HomeScreen({super.key});

  @override
  State<HomeScreen> createState() => HomeScreenState();
}

class HomeScreenState extends State<HomeScreen>
    with SingleTickerProviderStateMixin {
  late AnimationController _animationController;
  final EventService eventService = EventService();

  @override
  void initState() {
    super.initState();

    _animationController = AnimationController(
      vsync: this,
      duration: const Duration(milliseconds: 300),
      lowerBound: 0,
      upperBound: 1,
    );

    _animationController.forward();
  }

  Future<List<Event>> _fetchMyEvents() async {
    try {
      List<Event> events = await eventService.getMyEvents();

      return events;
    } catch (e) {
      // Handle error
      log.severe('Failed to fetch events: $e');
      return [];
    }
  }

  Future<List<Event>> _fetchRecommendedEvents() async {
    try {
      // TODO: Send Location data to the API
      List<Event> recommendedEvents = await eventService.getRecommendedEvents();

      debugPrint('recommendedEvents: $recommendedEvents');

      return recommendedEvents;
    } catch (e) {
      // Handle error
      log.severe('Failed to fetch events: $e');
      return [];
    }
  }

  @override
  void dispose() {
    _animationController.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return AnimatedBuilder(
      builder: (context, child) => SlideTransition(
        position: Tween(
          begin: const Offset(0, 0.3),
          end: const Offset(0, 0),
        ).animate(
          CurvedAnimation(
            parent: _animationController,
            curve: Curves.easeInOut,
          ),
        ),
        child: child,
      ),
      animation: _animationController,
      child: SingleChildScrollView(
        child: Column(
          children: [
            const SizedBox(
              height: 10,
            ),
            FutureBuilder<List<Event>>(
              future: _fetchMyEvents(),
              builder: (context, snapshot) {
                return Carousel(
                  text: "Mes événements",
                  isLoading:
                      snapshot.connectionState == ConnectionState.waiting,
                  items: snapshot.data != null
                      ? snapshot.data!
                          .map((event) =>
                              EventCard(event: event, hasJoinedEvent: true))
                          .toList()
                      : [],
                );
              },
            ),
            FutureBuilder<List<Event>>(
              future: _fetchRecommendedEvents(),
              builder: (context, snapshot) {
                return Carousel(
                  text: "Événements recommandés",
                  isLoading:
                      snapshot.connectionState == ConnectionState.waiting,
                  items: snapshot.data != null
                      ? snapshot.data!
                          .map((event) =>
                              EventCard(event: event, hasJoinedEvent: false))
                          .toList()
                      : [],
                );
              },
            ),
          ],
        ),
      ),
    );
  }
}
