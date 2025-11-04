package bekanius.mod;

import bekanius.mod.pkg.chakra.ChakraServer;
import net.fabricmc.api.ClientModInitializer;
import net.fabricmc.fabric.api.client.event.lifecycle.v1.ClientTickEvents;
import net.fabricmc.fabric.api.client.keybinding.v1.KeyBindingHelper;
import net.minecraft.client.option.KeyBinding;
import net.minecraft.client.util.InputUtil;
import net.minecraft.text.Text;
import org.lwjgl.glfw.GLFW;

public class ShinobiForgottenTalesClient implements ClientModInitializer {
    private static KeyBinding CHAKRA_RECHARGE_KEY;
    private static KeyBinding CHAKRA_USE_KEY;

    @Override
    public void onInitializeClient() {
        CHAKRA_RECHARGE_KEY = KeyBindingHelper.registerKeyBinding(new KeyBinding(
                "key.shinobi-forgotten-tales.charge_chakra",
                InputUtil.Type.KEYSYM,
                GLFW.GLFW_KEY_C,
                "category.shinobi-forgotten-tales.jutsus"
        ));

        CHAKRA_USE_KEY = KeyBindingHelper.registerKeyBinding(new KeyBinding(
                "key.shinobi-forgotten-tales.use_chakra",
                InputUtil.Type.KEYSYM,
                GLFW.GLFW_KEY_R,
                "category.shinobi-forgotten-tales.jutsus"
        ));

        ClientTickEvents.END_CLIENT_TICK.register(client -> {
            if (CHAKRA_RECHARGE_KEY.wasPressed()) {
                if (client.player != null && client.getServer() != null) {
                    client.player.sendMessage(Text.literal("Key 'C' was pressed."), false);
                    client.getServer().execute(() -> {
                        ChakraServer.rechargeChakra(client.player);
                    });
                }
            }

            if (CHAKRA_USE_KEY.wasPressed()) {
                if (client.player != null && client.getServer() != null) {
                    client.player.sendMessage(Text.literal("Key 'R' was pressed."), false);
                    client.getServer().execute(() -> {
                        ChakraServer.consumeChakra(client.player, 10);
                    });
                }
            }
        });
    }
}
