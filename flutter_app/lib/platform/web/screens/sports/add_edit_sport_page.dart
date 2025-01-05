import 'dart:ui';
import 'package:flutter/material.dart';
import 'package:squad_go/core/services/sport_service.dart';

class AddEditSportModal extends StatefulWidget {
  final Map<String, dynamic>? sport;
  final VoidCallback? onSportSaved;

  const AddEditSportModal({super.key, this.sport, this.onSportSaved});

  @override
  State<AddEditSportModal> createState() => _AddEditSportModalState();
}

class _AddEditSportModalState extends State<AddEditSportModal> {
  final _formKey = GlobalKey<FormState>();
  late String name;
  late String imageUrl;

  @override
  void initState() {
    super.initState();
    name = widget.sport?['name'] ?? '';
    imageUrl = widget.sport?['image_url'] ?? '';
  }

  Future<void> saveSport() async {
    if (!_formKey.currentState!.validate()) return;

    _formKey.currentState!.save();
    try {
      if (widget.sport == null) {
        await SportService.createSport({'name': name, 'image_url': imageUrl});
      } else {
        await SportService.updateSport(
          widget.sport!['id'],
          {'name': name, 'image_url': imageUrl},
        );
      }

      if (widget.onSportSaved != null) widget.onSportSaved!();
      Navigator.pop(context);
    } catch (e) {
      ScaffoldMessenger.of(context).showSnackBar(
        SnackBar(content: Text('Erreur: $e')),
      );
    }
  }

  @override
  Widget build(BuildContext context) {
    return BackdropFilter(
      filter: ImageFilter.blur(sigmaX: 10, sigmaY: 10),
      child: Dialog(
        shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(16)),
        backgroundColor: Colors.white,
        child: Container(
          padding: const EdgeInsets.all(16),
          width: MediaQuery.of(context).size.width * 0.8,
          constraints: const BoxConstraints(maxWidth: 400),
          child: Form(
            key: _formKey,
            child: Column(
              mainAxisSize: MainAxisSize.min,
              crossAxisAlignment: CrossAxisAlignment.stretch,
              children: [
                Align(
                  alignment: Alignment.topRight,
                  child: IconButton(
                    icon: const Icon(Icons.close),
                    onPressed: () => Navigator.pop(context),
                  ),
                ),
                Text(
                  widget.sport == null
                      ? 'Ajouter un Sport'
                      : 'Modifier le Sport',
                  textAlign: TextAlign.center,
                  style: const TextStyle(
                    fontSize: 18,
                    fontWeight: FontWeight.bold,
                  ),
                ),
                const SizedBox(height: 20),
                TextFormField(
                  initialValue: name,
                  decoration: InputDecoration(
                    labelText: 'Nom',
                    border: OutlineInputBorder(
                      borderRadius: BorderRadius.circular(8),
                    ),
                  ),
                  validator: (value) =>
                      value!.isEmpty ? 'Veuillez entrer un nom.' : null,
                  onSaved: (value) => name = value!,
                ),
                const SizedBox(height: 16),
                TextFormField(
                  initialValue: imageUrl,
                  decoration: InputDecoration(
                    labelText: 'URL de l\'Image',
                    border: OutlineInputBorder(
                      borderRadius: BorderRadius.circular(8),
                    ),
                  ),
                  validator: (value) =>
                      value!.isEmpty ? 'Veuillez entrer une URL valide.' : null,
                  onSaved: (value) => imageUrl = value!,
                ),
                const SizedBox(height: 24),
                Row(
                  mainAxisAlignment: MainAxisAlignment.spaceBetween,
                  children: [
                    TextButton(
                      onPressed: () => Navigator.pop(context),
                      style: TextButton.styleFrom(
                        foregroundColor: Colors.black,
                        backgroundColor: Colors.white,
                        side: const BorderSide(color: Colors.black),
                        padding: const EdgeInsets.symmetric(
                            horizontal: 24, vertical: 12),
                        shape: RoundedRectangleBorder(
                          borderRadius: BorderRadius.circular(8),
                        ),
                      ),
                      child: const Text('Annuler'),
                    ),
                    ElevatedButton(
                      onPressed: saveSport,
                      style: ElevatedButton.styleFrom(
                        backgroundColor: Colors.teal,
                        padding: const EdgeInsets.symmetric(
                            horizontal: 24, vertical: 12),
                        shape: RoundedRectangleBorder(
                          borderRadius: BorderRadius.circular(8),
                        ),
                      ),
                      child: const Text('Enregistrer'),
                    ),
                  ],
                ),
              ],
            ),
          ),
        ),
      ),
    );
  }
}
