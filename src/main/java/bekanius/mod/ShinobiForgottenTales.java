package bekanius.mod;

import bekanius.mod.pkg.chakra.ChakraServer;
import net.fabricmc.api.ModInitializer;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class ShinobiForgottenTales implements ModInitializer {
	public static final String MOD_ID = "shinobi-forgotten-tales";
	public static final Logger LOGGER = LoggerFactory.getLogger(MOD_ID);

	@Override
	public void onInitialize() {
		LOGGER.info("Inicializando o Chakra Server");
        ChakraServer.register();
	}
}