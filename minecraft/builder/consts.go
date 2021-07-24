package builder

import (
	"phoenixbuilder/minecraft/mctype"
	"github.com/lucasb-eyer/go-colorful"
)

var AirBlock = &mctype.ConstBlock{Name: "air", Data: 0}

var IronBlock = &mctype.ConstBlock{Name: "iron_block", Data: 0}

var ColorTable = []ColorBlock{
	{Block: &mctype.ConstBlock{Name: "dirt", Data: 0}, Color: colorful.Color{134, 96, 67}},
	{Block: &mctype.ConstBlock{Name: "cobblestone", Data: 0}, Color: colorful.Color{123, 123, 123}},
	{Block: &mctype.ConstBlock{Name: "bedrock", Data: 0}, Color: colorful.Color{84, 84, 84}},
	{Block: &mctype.ConstBlock{Name: "quartz_block", Data: 0}, Color: colorful.Color{237, 235, 228}},
	{Block: &mctype.ConstBlock{Name: "emerald_block", Data: 0}, Color: colorful.Color{81, 217, 117}},
	{Block: &mctype.ConstBlock{Name: "glowstone", Data: 0}, Color: colorful.Color{144, 118, 70}},
	{Block: &mctype.ConstBlock{Name: "gold_block", Data: 0}, Color: colorful.Color{249, 236, 79}},
	{Block: &mctype.ConstBlock{Name: "lapis_block", Data: 0}, Color: colorful.Color{39, 67, 138}},
	{Block: &mctype.ConstBlock{Name: "log", Data: 14}, Color: colorful.Color{207, 206, 201}},
	{Block: &mctype.ConstBlock{Name: "log", Data: 15}, Color: colorful.Color{87, 68, 27}},
	{Block: &mctype.ConstBlock{Name: "melon_block", Data: 0}, Color: colorful.Color{151, 154, 37}},
	{Block: &mctype.ConstBlock{Name: "netherrack", Data: 0}, Color: colorful.Color{111, 54, 53}},
	{Block: &mctype.ConstBlock{Name: "purpur_block", Data: 0}, Color: colorful.Color{170, 123, 170}},
	{Block: &mctype.ConstBlock{Name: "quartz_ore", Data: 0}, Color: colorful.Color{125, 85, 80}},
	{Block: &mctype.ConstBlock{Name: "redstone_block", Data: 0}, Color: colorful.Color{171, 28, 9}},
	{Block: &mctype.ConstBlock{Name: "sponge", Data: 0}, Color: colorful.Color{195, 196, 85}},
	{Block: &mctype.ConstBlock{Name: "slime", Data: 0}, Color: colorful.Color{121, 200, 101}},
	{Block: &mctype.ConstBlock{Name: "stained_hardened_clay", Data: 0}, Color: colorful.Color{210, 178, 162}},
	{Block: &mctype.ConstBlock{Name: "stained_hardened_clay", Data: 10}, Color: colorful.Color{120, 72, 88}},
	{Block: &mctype.ConstBlock{Name: "stained_hardened_clay", Data: 11}, Color: colorful.Color{75, 61, 93}},
	{Block: &mctype.ConstBlock{Name: "stained_hardened_clay", Data: 12}, Color: colorful.Color{78, 52, 37}},
	{Block: &mctype.ConstBlock{Name: "stained_hardened_clay", Data: 13}, Color: colorful.Color{104, 119, 54}},
	{Block: &mctype.ConstBlock{Name: "stained_hardened_clay", Data: 15}, Color: colorful.Color{37, 23, 17}},
	{Block: &mctype.ConstBlock{Name: "stained_hardened_clay", Data: 1}, Color: colorful.Color{163, 85, 39}},
	{Block: &mctype.ConstBlock{Name: "stained_hardened_clay", Data: 2}, Color: colorful.Color{151, 90, 111}},
	{Block: &mctype.ConstBlock{Name: "stained_hardened_clay", Data: 4}, Color: colorful.Color{188, 135, 37}},
	{Block: &mctype.ConstBlock{Name: "stained_hardened_clay", Data: 3}, Color: colorful.Color{114, 110, 140}},
	{Block: &mctype.ConstBlock{Name: "stained_hardened_clay", Data: 5}, Color: colorful.Color{104, 119, 54}},
	{Block: &mctype.ConstBlock{Name: "stained_hardened_clay", Data: 6}, Color: colorful.Color{164, 80, 80}},
	{Block: &mctype.ConstBlock{Name: "stained_hardened_clay", Data: 7}, Color: colorful.Color{58, 43, 37}},
	{Block: &mctype.ConstBlock{Name: "stained_hardened_clay", Data: 8}, Color: colorful.Color{136, 108, 99}},
	{Block: &mctype.ConstBlock{Name: "stone", Data: 1}, Color: colorful.Color{153, 114, 99}},
	{Block: &mctype.ConstBlock{Name: "stone", Data: 2}, Color: colorful.Color{159, 115, 98}},
	{Block: &mctype.ConstBlock{Name: "stone", Data: 5}, Color: colorful.Color{180, 180, 183}},
	{Block: &mctype.ConstBlock{Name: "stone", Data: 0}, Color: colorful.Color{125, 125, 125}},
	{Block: &mctype.ConstBlock{Name: "wool", Data: 0}, Color: colorful.Color{223, 223, 223}},
	{Block: &mctype.ConstBlock{Name: "wool", Data: 10}, Color: colorful.Color{125, 60, 180}},
	{Block: &mctype.ConstBlock{Name: "stained_hardened_clay", Data: 9}, Color: colorful.Color{87, 91, 91}},
	{Block: &mctype.ConstBlock{Name: "wool", Data: 11}, Color: colorful.Color{44, 54, 134}},
	{Block: &mctype.ConstBlock{Name: "wool", Data: 13}, Color: colorful.Color{54, 72, 28}},
	{Block: &mctype.ConstBlock{Name: "wool", Data: 14}, Color: colorful.Color{151, 52, 49}},
	{Block: &mctype.ConstBlock{Name: "wool", Data: 15}, Color: colorful.Color{37, 23, 17}},
	{Block: &mctype.ConstBlock{Name: "wool", Data: 12}, Color: colorful.Color{82, 52, 32}},
	{Block: &mctype.ConstBlock{Name: "wool", Data: 1}, Color: colorful.Color{219, 126, 64}},
	{Block: &mctype.ConstBlock{Name: "wool", Data: 2}, Color: colorful.Color{179, 79, 188}},
	{Block: &mctype.ConstBlock{Name: "wool", Data: 3}, Color: colorful.Color{107, 138, 201}},
	{Block: &mctype.ConstBlock{Name: "wool", Data: 4}, Color: colorful.Color{189, 177, 43}},
	{Block: &mctype.ConstBlock{Name: "wool", Data: 5}, Color: colorful.Color{69, 184, 59}},
	{Block: &mctype.ConstBlock{Name: "wool", Data: 6}, Color: colorful.Color{211, 141, 160}},
	{Block: &mctype.ConstBlock{Name: "wool", Data: 7}, Color: colorful.Color{65, 65, 65}},
	{Block: &mctype.ConstBlock{Name: "wool", Data: 8}, Color: colorful.Color{158, 164, 164}},
	{Block: &mctype.ConstBlock{Name: "wool", Data: 9}, Color: colorful.Color{47, 111, 138}},
	{Block: &mctype.ConstBlock{Name: "concrete", Data: 0}, Color: colorful.Color{207, 213, 214}},
	{Block: &mctype.ConstBlock{Name: "concrete", Data: 1}, Color: colorful.Color{224, 97, 0}},
	{Block: &mctype.ConstBlock{Name: "concrete", Data: 2}, Color: colorful.Color{169, 48, 159}},
	{Block: &mctype.ConstBlock{Name: "concrete", Data: 3}, Color: colorful.Color{35, 137, 198}},
	{Block: &mctype.ConstBlock{Name: "concrete", Data: 4}, Color: colorful.Color{241, 175, 21}},
	{Block: &mctype.ConstBlock{Name: "concrete", Data: 5}, Color: colorful.Color{94, 169, 25}},
	{Block: &mctype.ConstBlock{Name: "concrete", Data: 6}, Color: colorful.Color{213, 101, 142}},
	{Block: &mctype.ConstBlock{Name: "concrete", Data: 7}, Color: colorful.Color{55, 58, 62}},
	{Block: &mctype.ConstBlock{Name: "concrete", Data: 8}, Color: colorful.Color{125, 125, 115}},
	{Block: &mctype.ConstBlock{Name: "concrete", Data: 9}, Color: colorful.Color{21, 119, 136}},
	{Block: &mctype.ConstBlock{Name: "concrete", Data: 10}, Color: colorful.Color{100, 32, 156}},
	{Block: &mctype.ConstBlock{Name: "concrete", Data: 11}, Color: colorful.Color{45, 47, 143}},
	{Block: &mctype.ConstBlock{Name: "concrete", Data: 12}, Color: colorful.Color{96, 60, 32}},
	{Block: &mctype.ConstBlock{Name: "concrete", Data: 13}, Color: colorful.Color{73, 91, 36}},
	{Block: &mctype.ConstBlock{Name: "concrete", Data: 14}, Color: colorful.Color{142, 33, 33}},
	{Block: &mctype.ConstBlock{Name: "concrete", Data: 15}, Color: colorful.Color{8, 10, 15}},
	{Block: &mctype.ConstBlock{Name: "coal_block", Data: 0}, Color: colorful.Color{19, 19, 19}},
	{Block: &mctype.ConstBlock{Name: "diamond_block", Data: 0}, Color: colorful.Color{98, 219, 214}},
	{Block: &mctype.ConstBlock{Name: "dried_kelp_block", Data: 0}, Color: colorful.Color{50, 59, 39}},
	{Block: &mctype.ConstBlock{Name: "furnace", Data: 0}, Color: colorful.Color{96, 96, 96}},
	{Block: &mctype.ConstBlock{Name: "hay_block", Data: 0}, Color: colorful.Color{169, 140, 16}},
	{Block: &mctype.ConstBlock{Name: "iron_block", Data: 0}, Color: colorful.Color{219, 219, 219}},
	{Block: &mctype.ConstBlock{Name: "stripped_birch_log", Data: 0}, Color: colorful.Color{185, 161, 104}},
	{Block: &mctype.ConstBlock{Name: "stripped_acacia_log", Data: 0}, Color: colorful.Color{167, 92, 59}},
	{Block: &mctype.ConstBlock{Name: "stripped_jungle_log", Data: 0}, Color: colorful.Color{171, 134, 85}},
	{Block: &mctype.ConstBlock{Name: "stripped_oak_log", Data: 0}, Color: colorful.Color{164, 134, 81}},
	{Block: &mctype.ConstBlock{Name: "stripped_spruce_log", Data: 0}, Color: colorful.Color{106, 83, 48}},
	{Block: &mctype.ConstBlock{Name: "brick_block", Data: 0}, Color: colorful.Color{147, 100, 87}},
	{Block: &mctype.ConstBlock{Name: "clay", Data: 0}, Color: colorful.Color{159, 164, 177}},
	{Block: &mctype.ConstBlock{Name: "crafting_table", Data: 0}, Color: colorful.Color{107, 71, 43}},
	{Block: &mctype.ConstBlock{Name: "end_stone", Data: 0}, Color: colorful.Color{221, 224, 165}},
	{Block: &mctype.ConstBlock{Name: "red_glazed_terracotta", Data: 0}, Color: colorful.Color{182, 60, 53}},
	{Block: &mctype.ConstBlock{Name: "noteblock", Data: 0}, Color: colorful.Color{101, 68, 51}},
	{Block: &mctype.ConstBlock{Name: "sealantern", Data: 0}, Color: colorful.Color{172, 200, 190}},
	{Block: &mctype.ConstBlock{Name: "soul_sand", Data: 0}, Color: colorful.Color{85, 64, 52}},
	{Block: &mctype.ConstBlock{Name: "prismarine", Data: 0}, Color: colorful.Color{100, 152, 142}},
	{Block: &mctype.ConstBlock{Name: "pink_glazed_terracotta", Data: 0}, Color: colorful.Color{235, 155, 182}},
	{Block: &mctype.ConstBlock{Name: "purple_glazed_terracotta", Data: 0}, Color: colorful.Color{110, 48, 152}},
	{Block: &mctype.ConstBlock{Name: "magenta_glazed_terracotta", Data: 0}, Color: colorful.Color{208, 100, 192}},
	{Block: &mctype.ConstBlock{Name: "gray_glazed_terracotta", Data: 0}, Color: colorful.Color{83, 90, 94}},
	{Block: &mctype.ConstBlock{Name: "yellow_glazed_terracotta", Data: 0}, Color: colorful.Color{234, 192, 89}},
	{Block: &mctype.ConstBlock{Name: "blue_glazed_terracotta", Data: 0}, Color: colorful.Color{47, 65, 139}},
	{Block: &mctype.ConstBlock{Name: "obsidian", Data: 0}, Color: colorful.Color{20, 18, 30}},
	{Block: &mctype.ConstBlock{Name: "sponge", Data: 1}, Color: colorful.Color{160, 159, 63}},
	{Block: &mctype.ConstBlock{Name: "bone_block", Data: 0}, Color: colorful.Color{206, 201, 178}},
}

var FenceName="fence"
var PodzolName="podzol"

var BlockStr = []string{
	"air",
	"stone",
	"grass",
	"dirt",
	"cobblestone",
	"planks",
	"sapling",
	"bedrock",
	"flowing_water",
	"water",
	"flowing_lava",
	"lava",
	"sand",
	"gravel",
	"gold_ore",
	"iron_ore",
	"coal_ore",
	"log",
	"leaves",
	"sponge",
	"glass",
	"lapis_ore",
	"lapis_block",
	"dispenser",
	"sandstone",
	"noteblock",
	"bed",
	"golden_rail",
	"detector_rail",
	"sticky_piston",
	"cobweb",
	"tallgrass",
	"deadbush",
	"piston",
	"piston_head",
	"wool",
	"piston_extension",
	"dandelion",
	"poppy",
	"brown_mushroom",
	"red_mushroom",
	"gold_block",
	"iron_block",
	"double_stone_slab",
	"stone_slab",
	"brick_block",
	"tnt",
	"bookshelf",
	"mossy_cobblestone",
	"obsidian",
	"torch",
	"fire",
	"monster_spawner",
	"oak_stairs",
	"chest",
	"redstone_wire",
	"diamond_ore",
	"diamond_block",
	"crafting_table",
	"wheat",
	"farmland",
	"furnace",
	"lit_furnace",
	"standing_sign",
	"wooden_door",
	"ladder",
	"rail",
	"stone_stairs",
	"wall_sign",
	"lever",
	"stone_pressure_plate",
	"iron_door",
	"wooden_pressure_plate",
	"redstone_ore",
	"lit_redstone_ore",
	"unlit_redstone_torch",
	"redstone_torch",
	"stone_button",
	"snow_layer",
	"ice",
	"snow",
	"cactus",
	"clay",
	"reeds",
	"jukebox",
	"fence",
	"pumpkin",
	"netherrack",
	"soul_sand",
	"glowstone",
	"portal",
	"lit_pumpkin",
	"cake",
	"unpowered_repeater",
	"powered_repeater",
	"stained_glass",
	"trapdoor",
	"monster_egg",
	"stonebrick",
	"brown_mushroom_block",
	"red_mushroom_block",
	"iron_bars",
	"glass_pane",
	"melon_block",
	"pumpkin_stem",
	"melon_stem",
	"vine",
	"fence_gate",
	"brick_stairs",
	"stone_brick_stairs",
	"mycelium",
	"waterlily",
	"nether_brick",
	"nether_brick_fence",
	"nether_brick_stairs",
	"nether_wart",
	"enchanting_table",
	"brewing_stand",
	"cauldron",
	"end_portal",
	"end_portal_frame",
	"end_stone",
	"dragon_egg",
	"redstone_lamp",
	"lit_redstone_lamp",
	"double_wooden_slab",
	"wooden_slab",
	"cocoa",
	"sandstone_stairs",
	"emerald_ore",
	"ender_chest",
	"tripwire_hook",
	"tripwire",
	"emerald_block",
	"spruce_stairs",
	"birch_stairs",
	"jungle_stairs",
	"command_block",
	"beacon",
	"cobblestone_wall",
	"flower_pot",
	"carrots",
	"potatoes",
	"wooden_button",
	"skull",
	"anvil",
	"trapped_chest",
	"light_weighted_pressure_plate",
	"heavy_weighted_pressure_plate",
	"unpowered_comparator",
	"powered_comparator",
	"daylight_detector",
	"redstone_block",
	"quartz_ore",
	"hopper",
	"quartz_block",
	"quartz_stairs",
	"activator_rail",
	"dropper",
	"stained_hardened_clay",
	"stained_glass_pane",
	"leaves2",
	"log2",
	"acacia_stairs",
	"dark_oak_stairs",
	"slime",
	"barrier",
	"iron_trapdoor",
	"prismarine",
	"sealantern",
	"hay_block",
	"carpet",
	"hardened_clay",
	"coal_block",
	"packed_ice",
	"double_plant",
	"standing_banner",
	"wall_banner",
	"daylight_detector_inverted",
	"red_sandstone",
	"red_sandstone_stairs",
	"double_stone_slab2",
	"stone_slab2",
	"spruce_fence_gate",
	"birch_fence_gate",
	"jungle_fence_gate",
	"dark_oak_fence_gate",
	"acacia_fence_gate",
	"spruce_fence",
	"birch_fence",
	"jungle_fence",
	"dark_oak_fence",
	"acacia_fence_gate",
	"spruce_door",
	"birch_door",
	"jungle_door",
	"acacia_door",
	"dark_oak_door",
	"end_rod",
	"chorus_plant",
	"chorus_flower",
	"purpur_block",
	"purpur_pillar",
	"purpur_stairs",
	"purpur_double_slab",
	"purpur_slab",
	"end_bricks",
	"beetroots",
	"grass_path",
	"end_gateway",
	"repeating_command_block",
	"chain_command_block",
	"frosted_ice",
	"magma",
	"nether_wart_block",
	"red_nether_brick",
	"bone_block",
	"structure_void",
	"observer",
	"white_shulker_box",
	"orange_shulker_box",
	"magenta_shulker_box",
	"light_blue_shulker_box",
	"yellow_shulker_box",
	"lime_shulker_box",
	"pink_shulker_box",
	"gray_shulker_box",
	"light_gray_shulker_box",
	"cyan_shulker_box",
	"purple_shulker_box",
	"blue_shulker_box",
	"brown_shulker_box",
	"green_shulker_box",
	"red_shulker_box",
	"black_shulker_box",
	"white_glazed_terracotta",
	"orange_glazed_terracotta",
	"magenta_glazed_terracotta",
	"light_blue_glazed_terracotta",
	"yellow_glazed_terracotta",
	"lime_glazed_terracotta",
	"pink_glazed_terracotta",
	"gray_glazed_terracotta",
	"light_gray_glazed_terracotta",
	"cyan_glazed_terracotta",
	"purple_glazed_terracotta",
	"blue_glazed_terracotta",
	"brown_glazed_terracotta",
	"green_glazed_terracotta",
	"red_glazed_terracotta",
	"black_glazed_terracotta",
	"concrete",
	"concrete_powder",
	"null",
	"null",
	"structure_block",
}
