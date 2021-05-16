class Tag {
  public Type: string | undefined;
  public Id: number | undefined;
  public Value: string | undefined;
}

export class BoardGame {
  public Type: string | undefined;
  public ObjectId: number | undefined;
  public MinPlayers: number | undefined;
  public MaxPlayers: number | undefined;
  public Thumb: string | undefined;
  public Image: string | undefined;
  public Description: string | undefined;
  public Name: string | undefined;
  public Names: string[] | undefined;
  public Tags: Tag[] | undefined;
  /*
{
    "Type": "boardgame",
    "ObjectId": 266192,
    "MinPlayers": 1,
    "MaxPlayers": 5,
    "Image": "https://cf.geekdo-images.com/yLZJCVLlIx4c7eJEWUNJ7w__thumb/img/VNToqgS2-pOGU6MuvIkMPKn_y-s=/fit-in/200x150/filters:strip_icc()/pic4458123.jpg",
    "Thumb": "https://cf.geekdo-images.com/yLZJCVLlIx4c7eJEWUNJ7w__thumb/img/VNToqgS2-pOGU6MuvIkMPKn_y-s=/fit-in/200x150/filters:strip_icc()/pic4458123.jpg",
    "Description": "Wingspan is&nbsp;a competitive, medium-weight, card-driven, engine-building board game from Stonemaier Games. It's designed by Elizabeth Hargrave and features over 170 birds illustrated by Natalia Rojas and Ana Maria Martinez.&#10;&#10;You are bird enthusiasts&mdash;researchers, bird watchers, ornithologists, and collectors&mdash;seeking to discover and attract the best birds to your network of wildlife preserves. Each bird extends a chain of powerful combinations in one of your habitats (actions). These habitats  focus on several key aspects of growth:&#10;&#10;&#10;     Gain food tokens via custom dice in a birdfeeder dice tower&#10;     Lay eggs using egg miniatures in a variety of colors&#10;     Draw from hundreds of unique bird cards and play them&#10;&#10;&#10;The winner is the player with the most points after 4 rounds.&#10;&#10;If you enjoy Terraforming Mars and Gizmos, we think this game will take flight at your table.&#10;&#10;&mdash;description from the publisher&#10;&#10;From the 7th printing on, the base game box includes Wingspan: Swift-Start Promo Pack.&#10;&#10;",
    "Tags": [
        {
            "Type": "boardgamecategory",
            "Id": 1089,
            "Value": "Animals"
        },
        {
            "Type": "boardgamecategory",
            "Id": 1002,
            "Value": "Card Game"
        },
        {
            "Type": "boardgamecategory",
            "Id": 1094,
            "Value": "Educational"
        },
        {
            "Type": "boardgamemechanic",
            "Id": 2041,
            "Value": "Card Drafting"
        },
        {
            "Type": "boardgamemechanic",
            "Id": 2072,
            "Value": "Dice Rolling"
        },
        {
            "Type": "boardgamemechanic",
            "Id": 2984,
            "Value": "Drafting"
        },
        {
            "Type": "boardgamemechanic",
            "Id": 2875,
            "Value": "End Game Bonuses"
        },
        {
            "Type": "boardgamemechanic",
            "Id": 2040,
            "Value": "Hand Management"
        },
        {
            "Type": "boardgamemechanic",
            "Id": 2004,
            "Value": "Set Collection"
        },
        {
            "Type": "boardgamemechanic",
            "Id": 2819,
            "Value": "Solo / Solitaire Game"
        },
        {
            "Type": "boardgamefamily",
            "Id": 45672,
            "Value": "Animals: Birds"
        },
        {
            "Type": "boardgamefamily",
            "Id": 7657,
            "Value": "Animals: Crows / Ravens / Magpies"
        },
        {
            "Type": "boardgamefamily",
            "Id": 7306,
            "Value": "Animals: Ducks"
        },
        {
            "Type": "boardgamefamily",
            "Id": 18887,
            "Value": "Animals: Geese"
        },
        {
            "Type": "boardgamefamily",
            "Id": 18588,
            "Value": "Animals: Owls"
        },
        {
            "Type": "boardgamefamily",
            "Id": 17963,
            "Value": "Animals: Turkeys"
        },
        {
            "Type": "boardgamefamily",
            "Id": 66553,
            "Value": "Components: Control Boards"
        },
        {
            "Type": "boardgamefamily",
            "Id": 65328,
            "Value": "Components: Dice with Icons"
        },
        {
            "Type": "boardgamefamily",
            "Id": 66772,
            "Value": "Components: Drop Tower"
        },
        {
            "Type": "boardgamefamily",
            "Id": 48877,
            "Value": "Components: Game Trayz Inside"
        },
        {
            "Type": "boardgamefamily",
            "Id": 58267,
            "Value": "Game: Wingspan"
        },
        {
            "Type": "boardgamefamily",
            "Id": 27646,
            "Value": "Mechanism: Tableau Building"
        },
        {
            "Type": "boardgamefamily",
            "Id": 5666,
            "Value": "Players: Games with Solitaire Rules"
        },
        {
            "Type": "boardgamefamily",
            "Id": 59442,
            "Value": "Theme: Biology"
        },
        {
            "Type": "boardgamefamily",
            "Id": 48871,
            "Value": "Theme: Nature"
        },
        {
            "Type": "boardgameexpansion",
            "Id": 290448,
            "Value": "Wingspan: European Expansion"
        },
        {
            "Type": "boardgameexpansion",
            "Id": 300580,
            "Value": "Wingspan: Oceania Expansion"
        },
        {
            "Type": "boardgameexpansion",
            "Id": 290837,
            "Value": "Wingspan: Swift-Start Promo Pack"
        },
        {
            "Type": "boardgamedesigner",
            "Id": 111338,
            "Value": "Elizabeth Hargrave"
        },
        {
            "Type": "boardgameartist",
            "Id": 113749,
            "Value": "Ana Maria Martinez Jaramillo"
        },
        {
            "Type": "boardgameartist",
            "Id": 113748,
            "Value": "Natalia Rojas"
        },
        {
            "Type": "boardgameartist",
            "Id": 71164,
            "Value": "Beth Sobel"
        },
        {
            "Type": "boardgamepublisher",
            "Id": 23202,
            "Value": "Stonemaier Games"
        },
        {
            "Type": "boardgamepublisher",
            "Id": 267,
            "Value": "999 Games"
        },
        {
            "Type": "boardgamepublisher",
            "Id": 38809,
            "Value": "Angry Lion Games"
        },
        {
            "Type": "boardgamepublisher",
            "Id": 3475,
            "Value": "Arclight"
        },
        {
            "Type": "boardgamepublisher",
            "Id": 7162,
            "Value": "Brain Games"
        },
        {
            "Type": "boardgamepublisher",
            "Id": 6194,
            "Value": "Delta Vision Publishing"
        },
        {
            "Type": "boardgamepublisher",
            "Id": 40415,
            "Value": "Divercentro"
        },
        {
            "Type": "boardgamepublisher",
            "Id": 22380,
            "Value": "Feuerland Spiele"
        },
        {
            "Type": "boardgamepublisher",
            "Id": 4785,
            "Value": "Ghenos Games"
        },
        {
            "Type": "boardgamepublisher",
            "Id": 46884,
            "Value": "Green Elephant Games"
        },
        {
            "Type": "boardgamepublisher",
            "Id": 6214,
            "Value": "Kaissa Chess & Games"
        },
        {
            "Type": "boardgamepublisher",
            "Id": 8291,
            "Value": "Korea Boardgames Co., Ltd."
        },
        {
            "Type": "boardgamepublisher",
            "Id": 3218,
            "Value": "Lautapelit.fi"
        },
        {
            "Type": "boardgamepublisher",
            "Id": 34801,
            "Value": "Lavka Games"
        },
        {
            "Type": "boardgamepublisher",
            "Id": 29242,
            "Value": "Ludofy Creative"
        },
        {
            "Type": "boardgamepublisher",
            "Id": 30677,
            "Value": "Maldito Games"
        },
        {
            "Type": "boardgamepublisher",
            "Id": 5400,
            "Value": "Matagot"
        },
        {
            "Type": "boardgamepublisher",
            "Id": 7992,
            "Value": "MINDOK"
        },
        {
            "Type": "boardgamepublisher",
            "Id": 32395,
            "Value": "NeoTroy Games"
        },
        {
            "Type": "boardgamepublisher",
            "Id": 7466,
            "Value": "Rebel Sp. z o.o."
        },
        {
            "Type": "boardgamepublisher",
            "Id": 44241,
            "Value": "Regatul Jocurilor"
        },
        {
            "Type": "boardgamepublisher",
            "Id": 33998,
            "Value": "Siam Board Games"
        },
        {
            "Type": "boardgamepublisher",
            "Id": 36763,
            "Value": "Surfin' Meeple China"
        },
        {
            "Type": "boardgamepublisher",
            "Id": 44209,
            "Value": "Ігромаг"
        }
    ],
    "Names": [
        "Wingspan",
        "Fesztáv",
        "Flügelschlag",
        "Na křídlech",
        "Na skrzydłach",
        "Spārnotie",
        "Sparnuotieji",
        "Tiivulised",
        "ΑΝΟΙΓΜΑ ΦΤΕΡΩΝ",
        "Крила",
        "Крылья",
        "ปีกปักษา",
        "ウイングスパン",
        "展翅翱翔",
        "윙스팬"
    ]
}
  */
}
