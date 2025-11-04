package bekanius.mod.pkg.chakra;

import net.fabricmc.fabric.api.networking.v1.ServerPlayConnectionEvents;
import net.minecraft.entity.player.PlayerEntity;
import net.minecraft.text.Text;

import java.util.HashMap;
import java.util.Map;

public class ChakraServer {
    private static final Map<PlayerEntity, Integer> playerChakra = new HashMap<>();


    public static void register() {
        ServerPlayConnectionEvents.JOIN.register((handler, sender, server) -> {
            playerChakra.put(handler.player, 100);
            handler.player.sendMessage(Text.literal("Chakra do jogador: " + handler.player.getName() + " setado."), false);
        });
    }

    public static boolean consumeChakra(PlayerEntity plr, int amount) {
        int currentChakra = playerChakra.getOrDefault(plr, 100);
        if (currentChakra >= amount) {
            playerChakra.put(plr, currentChakra - amount);
            plr.sendMessage(Text.literal("Consumindo " + amount + " chakra, restaram: " + String.valueOf(currentChakra - amount)));
            return true;
        }
        return false;
    }

    public static boolean rechargeChakra(PlayerEntity plr) {
        int currentChakra = playerChakra.getOrDefault(plr, 100);
        if (currentChakra < 100) {
            playerChakra.put(plr, 100);
            plr.sendMessage(Text.literal("Chakra recarregado."));
            return true;
        }
        return false;
    }
}
