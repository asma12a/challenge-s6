import 'package:flutter/material.dart';
import 'package:carousel_slider/carousel_slider.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';

class Carousel extends StatefulWidget {
  const Carousel({
    super.key,
    required this.items,
    required this.text,
    this.isLoading = false,
  });

  final List<Widget> items;
  final String text;
  final bool isLoading;

  @override
  State<Carousel> createState() => CarouselState();
}

class CarouselState extends State<Carousel> {
  @override
  Widget build(BuildContext context) {
    final translate = AppLocalizations.of(context);
    final translate = AppLocalizations.of(context);
    return LayoutBuilder(
      builder: (context, constraints) {
        return Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Padding(
              padding: const EdgeInsets.only(left: 40),
              child: Text(
                widget.text,
                style: Theme.of(context).textTheme.titleLarge!,
              ),
            ),
            const SizedBox(
              height: 10,
            ),
            if (widget.isLoading)
              Padding(
                padding: const EdgeInsets.only(left: 40),
                child: const Center(
                  child: CircularProgressIndicator(),
                ),
              ),
            if (!widget.isLoading && widget.items.isEmpty)
              Padding(
                padding: const EdgeInsets.only(left: 40),
                child: Text(translate?.no_events_to_display ?? 'Aucun événement à afficher'),
              ),
            Padding(
              padding: const EdgeInsets.only(left: 20),
              child: CarouselSlider(
                options: CarouselOptions(
                  enableInfiniteScroll: false,
                  padEnds: false,
                  height: widget.items.isEmpty ? 20 : null,
                ),
                items: widget.items,
              ),
            ),
            const SizedBox(
              height: 20,
            ),
          ],
        );
      },
    );
  }
}
