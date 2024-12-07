import 'package:flutter/material.dart';
import 'package:squad_go/models/event.dart';
import 'package:squad_go/models/sport.dart';
import 'package:squad_go/widgets/carousel.dart';
import 'package:squad_go/widgets/event_card.dart';

class HomeScreen extends StatefulWidget {
  const HomeScreen({super.key});

  @override
  State<HomeScreen> createState() => HomeScreenState();
}

class HomeScreenState extends State<HomeScreen>
    with SingleTickerProviderStateMixin {
  late AnimationController _animationController;

  // TODO: get events from the API (my events and recommended events)

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
            Carousel(
              text: "Mes événements",
              items: [
                EventCard(
                  event: Event(
                    id: "id",
                    name:
                        "Event 1 - Football Event 1 - Football Event 1 - Football Event 1 - Football Event 1 - Football Event 1 - Football Event 1 - Football ",
                    address:
                        "16, Rue de la rue 16, Rue de la rue 16, Rue de la rue 16, Rue de la rue 16, Rue de la rue 16, Rue de la rue 16, Rue de la rue 16, Rue de la rue 16, Rue de la rue ",
                    date: "2022-01-01",
                    sport: Sport(
                      id: "id",
                      name: SportName.football,
                      type: SportType.team,
                      color: Colors.blue,
                    ),
                  ),
                  hasJoinedEvent: true,
                ),
                EventCard(
                  event: Event(
                    id: "id2",
                    name: "Event 2",
                    address: "Rue de la rue",
                    date: "2022-01-01",
                    type: EventType.training,
                    sport: Sport(
                      id: "id2",
                      name: SportName.basketball,
                      type: SportType.team,
                    ),
                  ),
                  hasJoinedEvent: true,
                ),
                EventCard(
                  event: Event(
                    id: "id3",
                    name: "Event 2",
                    address: "Rue de la rue",
                    date: "2022-01-01",
                    sport: Sport(
                      id: "id3",
                      name: SportName.tennis,
                      type: SportType.individual,
                    ),
                  ),
                ),
                EventCard(
                  event: Event(
                    id: "id4",
                    name: "Event 2",
                    address: "Rue de la rue",
                    date: "2022-01-01",
                    sport: Sport(
                      id: "id4",
                      name: SportName.running,
                      type: SportType.individual,
                    ),
                  ),
                ),
              ],
            ),
            Carousel(
              text: "Évenements recommandés",
              items: [],
            ),
          ],
        ),
      ),
    );
  }
}
