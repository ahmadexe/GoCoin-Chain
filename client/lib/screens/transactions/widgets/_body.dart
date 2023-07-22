part of '../transactions.dart';

class _Body extends StatelessWidget {
  const _Body();

  static final GlobalKey<FormBuilderState> _formKey =
      GlobalKey<FormBuilderState>();

  @override
  Widget build(BuildContext context) {
    final appMedia = MediaQuery.sizeOf(context);

    return SingleChildScrollView(
      padding: EdgeInsets.all(appMedia.width * 0.05),
      child: FormBuilder(
        key: _formKey,
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            const Text(
              'Wallet Information',
              style: TextStyle(
                fontSize: 24,
                fontWeight: FontWeight.w600,
              ),
            ),
            const SizedBox(
              height: 30,
            ),
            const Text(
              'Public Key',
              style: TextStyle(
                fontSize: 18,
                fontWeight: FontWeight.w600,
              ),
            ),
            const SizedBox(
              height: 10,
            ),
            FormBuilderTextField(
              name: 'public-key',
              decoration: InputDecoration(
                hintText: 'Public Key',
                prefixIcon: const Icon(Icons.key),
                border: OutlineInputBorder(
                  borderRadius: BorderRadius.circular(10),
                ),
              ),
            ),
            const SizedBox(
              height: 20,
            ),
            const Text(
              'Private Key',
              style: TextStyle(
                fontSize: 18,
                fontWeight: FontWeight.w600,
              ),
            ),
            const SizedBox(
              height: 10,
            ),
            FormBuilderTextField(
              name: 'private-key',
              decoration: InputDecoration(
                hintText: 'Private Key',
                prefixIcon: const Icon(Icons.key),
                border: OutlineInputBorder(
                  borderRadius: BorderRadius.circular(10),
                ),
              ),
            ),
            const SizedBox(
              height: 20,
            ),
            const Text(
              'Blockchain Address',
              style: TextStyle(
                fontSize: 18,
                fontWeight: FontWeight.w600,
              ),
            ),
            const SizedBox(
              height: 10,
            ),
            FormBuilderTextField(
              name: 'chain-address',
              decoration: InputDecoration(
                hintText: 'Blockchain Address',
                prefixIcon: const Icon(Icons.key),
                border: OutlineInputBorder(
                  borderRadius: BorderRadius.circular(10),
                ),
              ),
            ),
            const SizedBox(
              height: 30,
            ),
            const Text(
              'Send Coins',
              style: TextStyle(
                fontSize: 24,
                fontWeight: FontWeight.w600,
              ),
            ),
            const SizedBox(
              height: 20,
            ),
            const Text(
              'Blockchain Address',
              style: TextStyle(
                fontSize: 18,
                fontWeight: FontWeight.w600,
              ),
            ),
            const SizedBox(
              height: 10,
            ),
            FormBuilderTextField(
              name: 'receiver-address',
              decoration: InputDecoration(
                hintText: 'Receivers Address',
                prefixIcon: const Icon(Icons.key),
                border: OutlineInputBorder(
                  borderRadius: BorderRadius.circular(10),
                ),
              ),
            ),
            const SizedBox(
              height: 20,
            ),
            const Text(
              'Blockchain Address',
              style: TextStyle(
                fontSize: 18,
                fontWeight: FontWeight.w600,
              ),
            ),
            const SizedBox(
              height: 10,
            ),
            FormBuilderTextField(
              name: 'ammount',
              decoration: InputDecoration(
                hintText: 'Ammount',
                prefixIcon: const Icon(Icons.key),
                border: OutlineInputBorder(
                  borderRadius: BorderRadius.circular(10),
                ),
              ),
            ),
          ],
        ),
      ),
    );
  }
}
