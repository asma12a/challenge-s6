import 'package:flutter/material.dart';
import 'package:flutter_app/models/event.dart';
import 'package:flutter_app/widgets/carousel.dart';
import 'package:flutter_app/widgets/event_card.dart';

class HomeScreen extends StatefulWidget {
  const HomeScreen({super.key});

  @override
  State<HomeScreen> createState() => HomeScreenState();
}

class HomeScreenState extends State<HomeScreen>
    with SingleTickerProviderStateMixin {
  late AnimationController _animationController;

  @override
  void initState() {
    // TODO: implement initState
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
    // TODO: implement dispose
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
            Carousel(
              text: "Mes événements",
              items: [
                EventCard(
                  event: Event(
                    id: "id",
                    name: "Event 1",
                    address: "Rue de la rue",
                    date: "2022-01-01",
                    sport: "football",
                  ),
                ),
                EventCard(
                  event: Event(
                    id: "id2",
                    name: "Event 2",
                    address: "Rue de la rue",
                    date: "2022-01-01",
                    sport: "football",
                  ),
                ),
              ],
            ),
            SizedBox(
              height: 100,
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
