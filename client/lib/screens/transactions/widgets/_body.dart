part of '../transactions.dart';

class _Body extends StatefulWidget {
  const _Body();

  @override
  State<_Body> createState() => _BodyState();
}

class _BodyState extends State<_Body> {
  @override
  void initState() {
    super.initState();
    final walletBloc = WalletBloc.get(context, false);
    walletBloc.add(GetWalletDetails());
  }

  static final GlobalKey<FormBuilderState> _formKey =
      GlobalKey<FormBuilderState>();
  @override
  Widget build(BuildContext context) {
    final appMedia = MediaQuery.sizeOf(context);
    final transactionBloc = TransactionBloc.get(context, false);

    return SingleChildScrollView(
      padding: EdgeInsets.all(appMedia.width * 0.05),
      child: BlocBuilder<WalletBloc, WalletState>(
        builder: (context, state) {
          if (state is WalletLoading || state is WalletInitial) {
            return const Center(
              child: CircularProgressIndicator(),
            );
          } else if (state is WalletLoaded) {
            final wallet = state.wallet!;

            return FormBuilder(
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
                    initialValue: wallet.publicKey,
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
                    initialValue: wallet.privateKey,
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
                    initialValue: wallet.blockchainAddress,
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
                      prefixIcon: const Icon(Icons.security_sharp),
                      border: OutlineInputBorder(
                        borderRadius: BorderRadius.circular(10),
                      ),
                    ),
                  ),
                  const SizedBox(
                    height: 20,
                  ),
                  const Text(
                    'Ammount',
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
                      prefixIcon: const Icon(Icons.home_work_sharp),
                      border: OutlineInputBorder(
                        borderRadius: BorderRadius.circular(10),
                      ),
                    ),
                  ),
                  const SizedBox(
                    height: 20,
                  ),
                  BlocBuilder<TransactionBloc, TransactionState>(
                    builder: (context, transactionState) {
                      if (transactionState is TransactionLoading) {
                        return const Center(
                          child: CircularProgressIndicator(),
                        );
                      }
                      return SizedBox(
                        width: double.infinity,
                        height: 60,
                        child: ElevatedButton(
                          onPressed: () {
                            final form = _formKey.currentState!;
                            form.save();

                            final isValid = form.saveAndValidate();
                            if (!isValid) return;

                            final data = form.value;
                            final receiverAddress = data['receiver-address'];
                            final ammount = data['ammount'];

                            final transaction = Transaction(
                              senderPublicKey: wallet.publicKey,
                              senderPrivateKey: wallet.privateKey,
                              senderBlockchainAddress: wallet.privateKey,
                              recipientBlockchainAddress: receiverAddress,
                              value: double.parse(ammount),
                            );

                            transactionBloc.add(
                              CreateTransaction(transaction: transaction),
                            );
                          },
                          style: ElevatedButton.styleFrom(
                            backgroundColor: Colors.black,
                            foregroundColor: Colors.white,
                            shape: RoundedRectangleBorder(
                              borderRadius: BorderRadius.circular(10),
                            ),
                          ),
                          child: const Text(
                            'Send',
                            style: TextStyle(fontSize: 18),
                          ),
                        ),
                      );
                    },
                  ),
                ],
              ),
            );
          } else {
            return Center(
              child: Text(state.message!),
            );
          }
        },
      ),
    );
  }
}
