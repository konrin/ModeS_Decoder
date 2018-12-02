package main

import (
	"adsb_decoder"
	"fmt"
	"time"
)

func main() {

	decoder := adsbdecoder.NewDecoder(adsbdecoder.CACHE_TTL)

	_ = decoder

	for _, m := range messages {
		msg := adsbdecoder.NewMessage(m, time.Now())

		err := decoder.Decode(msg)
		if err != nil {
			println(err.Error())
			continue
		}

		fmt.Printf("%d, %s, %s, %d, (%f, %f), %f, %f, %d, %s, %d\n",
			msg.DF,
			msg.ICAO,
			msg.Callsign,
			msg.Category,
			msg.Lat,
			msg.Lon,
			msg.Speed,
			msg.Track,
			msg.Rocd,
			msg.Tag,
			msg.Altitude,
		)
		//fmt.Printf("%+v\n", msg)
	}
}

var messages = []string{
	"8d4850b558b3f7b4cabb5faf2a45",
	"8d71c046ea40885cd73c087c186b",
	"8d3c4b2e58b5046644b03e2c81fa",
	"8d71c04658ab00f3a73c5ef9daf9",
	"8d47340a990da587b804c8b2cba1",
	"5d71c046fdcc43",
	"8d4850b59909a99ed008c28522d8",
	"5d7892181dc6e1",
	"8d47340a60b9802ab9e406d002b0",
	"5d3c4b2eba4b44",
	"8d4850b59909a99ed008c28522d8",
	"8d3c4b2e58b5010434f5e85736ee",
	"8d3c4b2ee990b4247138005493a1",
	"8d71c0469909f5855804ca1a8140",
	"8d71c046252c1335cf0820cc7b1d",
	"8d71c046252c1335cf0820cc7b1d",
	"8d3c4b2ef801000000392861ac04",
	"8d4850b59909a99ed008c28522d8",
	"8d3c4b2e58b5046644b07e2f029a",
	"8d4850b5ea447864013c0866d022",
	"8d3c4b2e991206006808ca1e323b",
	"8d47340a60b9802aafe3e1f10b92",
	"5d47340a29f4f9",
	"5d47340a29f4f9",
	"5d4850b5e506cd",
	"5d47340a29f4f9",
	"8d71c04658ab0455f0f71de10552",
	"5d47340a29f4f9",
	"8d71c04658ab0455f0f71de10552",
	"5d47340a29f4f9",
	"8d3c4b2e991207006808ca1d38f5",
	"8d789218ea502864af5c089a8c40",
	"8d44003f9901ec848807ce28c9f9",
	"8d71c04658ab0455ecf735640392",
	"8d71c04658ab0455ecf735640392",
	"8d3c4b2e58b5046644b0aa2b5fd4",
	"8d4850b5e1163e00000000927cfd",
	"8d3c4b2e58b5046644b0aa2b5fd4",
	"8d4850b5e1163e00000000927cfd",
	"5d71c046fdcc43",
	"8d4244449909c9849804c79cf4a6",
	"5d424444b2b286",
	"8d47340a60b9802aa3e3b5aba310",
	"a0001690c4662330aa000028285e",
	"a0001690c4662330aa000028285e",
	"8d71c04658ab0455e8f74951b8c4",
	"a0001690ffd3ff40e004f73f65ee",
	"a000169099aa49357fefffbfbb51",
	"a0001690c4662330aa000028285e",
	"a00015308054373ec004eed52111",
	"a0001690ffd3ff40e004f73f65ee",
	"a0001530c0780030a4000029a8b5",
	"a0001690ffd3ff40e004f73f65ee",
	"a00015308054373ec004eed52111",
	"a000169099aa49357fefffbfbb51",
	"8d4850b5ea447864013c0866d022",
	"a00015308054373ec004eed52111",
	"a0001690c4662330aa000028285e",
	"a0001690ffd3ff40e004f73f65ee",
	"a0001690ffd3ff40e004f73f65ee",
	"5d3c4b2eba4b44",
	"5d4850b5e506cd",
	"8d7892189911f40cb004bd2f1c12",
	"8d4850b5ea447864013c0866d022",
	"a00015308054373ec004eed52111",
	"a0001690c4662330aa000028285e",
	"a0001690ffd3ff40e004f73f65ee",
	"a0001690ffd3ff40e004f73f65ee",
	"5d3c4b2eba4b44",
	"5d4850b5e506cd",
	"8d7892189911f40cb004bd2f1c12",
	"5d4850b5e506cd",
	"8d3c4b2e991207006808ca1d38f5",
	"8d71c0469909f5855808ca52db40",
	"8d44003f9901ec848807ce28c9f9",
	"8d3c4b2e2510c237cb0820f55865",
	"5d47340a29f4f9",
	"5d3c4b2eba4b44",
	"5d3c4b2eba4b44",
	"5d3c4b2eba4b44",
	"a800163effb5573d7ffce1565518",
	"a000163fa56a0f30bfb402d94f30",
	"a800163ec4600030aa0000824f14",
	"a800163ec4600030aa0000824f14",
	"5d424444b2b286",
	"8d3c4b2e58b5046644b0e028b0c3",
	"5d424444b2b286",
	"8d4244449909c9849804c79cf4a6",
	"5d44003f27cf64",
	"8d47340a990da587b004c8dc69a9",
	"8d4850b558b3f04f0cffabb6327a",
	"8d3c4b2e991207006808ca1d38f5",
	"8d47340a60b98790479b35acf657",
	"8d3c4b2ebf1418000000006eb9b3",
	"8d3c4b2e58b5046646b0f83388f5",
	"5d4850b5e506cd",
	"8d4850b5ea447864013c0866d022",
	"8d71c0469909f5855808ca52db40",
	"8d3c4b2e991207006808ca1d38f5",
	"8d71c04658ab00f3853d22abe980",
	"5d71c046fdcc43",
	"8d3c4b2e58b5046646b10c38198b",
	"8d3c4b2e58b5046646b10c38198b",
	"8d4850b558b3f7b3f2bc105c4165",
	"8d3c4b2ef801000000392861ac04",
	"8d71c04658ab00f3833d3687c8e8",
	"8d4850b5232da4b8df6820be8bb1",
	"5d3c4b2eba4b44",
	"5d4850b5e506cd",
	"8d47340a60b98790399b07fd50c1",
	"5d4850b5e506cd",
	"5d4850b5e506cd",
	"5d4850b5e506cd",
	"8d4850b558b3f7b3e0bc209b8c9f",
	"8d3c4b2ee990b4247138005493a1",
	"5d4850b5e506cd",
	"8d3c4b2e58b5010436f6b25d8723",
	"8d71c0469909f5855808ca52db40",
	"8d3c4b2e58b5010436f6b25d8723",
	"8d3c4b2e991207006808ca1d38f5",
	"5d4850b5e506cd",
	"8d47340a990da587b004c8dc69a9",
	"8d71c0469909f5857808ca17a769",
	"8d4850b558b3f04eaefffad5889e",
	"5d4850b5e506cd",
	"8d44003f20055076ce08201bc6d3",
	"8d44003f60b5045b3090ccef893e",
	"8d3c4b2e991207006808ca1d38f5",
	"8d3c4b2ee990b4247138005493a1",
	"8d71c04658ab0455ccf7f325de36",
	"5d44003f27cf64",
	"8d71c046f82100020049b891fb5a",
	"8d3c4b2e58b5010436f6fe5e4c19",
	"8d47340a60b98790239abea2b016",
	"8d3c4b2e991207006808ca1d38f5",
	"5d71c046fdcc43",
	"8d71c0469909f5857808ca17a769",
	"8d47340aea466858011c08981ac9",
	"8d71c04658ab00f36d3db318d94f",
	"8d3c4b2e58b5046646b1903fb7c9",
	"8d3c4b2e991207006808ca1d38f5",
	"8d3c4b2e991207006808ca1d38f5",
	"5d44003f27cf64",
	"8d47340a990da587b004c8dc69a9",
	"8d4850b558b3f04e4300555c0082",
	"8d3c4b2e58b5046646b1ae3ec2e0",
	"5d44003f27cf64",
	"8d3c4b2e991207006808ca1d38f5",
	"8d4850b59909a99ed008c28522d8",
	"8d71c046e1028600000000f160a0",
	"8d44003f9901ec846007ce63eb27",
	"8d3c4b2e2510c237cb0820f55865",
	"8d3c4b2e58b5010436f74fa99238",
	"8d71c04658ab00f3653de68ad4c0",
	"8d71c0469909f5857808ca17a769",
	"8d3c4b2ee990b4247138005493a1",
	"8d3c4b2e991207004808ca5844dc",
	"8d71c0469909f5857808ca17a769",
	"5d44003f27cf64",
	"5d3c4b2eba4b44",
	"8d71c046f82100020049b891fb5a",
	"8d3c4b2e58b5046646b1f23de902",
	"8d3c4b2e991207004808ca5844dc",
	"5d4850b5e506cd",
	"8d3c4b2ef801000000392861ac04",
	"8d71c0469909f5857808ca17a769",
	"8d47340a990da587b008c89433a9",
	"8d71c04658ab0455aef89e5640cc",
	"8d47340a60b9878fff9a4170da76",
	"8d3c4b2ee990b4247138005493a1",
	"8d44003f60b5045b12918343733d",
	"8d71c0469909f5857808ca17a769",
	"8d71c046ea40885cd73c087c186b",
	"8d3c4b2e58b5010436f7b453ae8e",
	"8d3c4b2e991207004808ca5844dc",
	"8d71c0469909f5857808ca17a769",
	"8d71c0469909f5857808ca17a769",
	"8d4850b558b3f04db700ca17d4e5",
	"5d71c046fdcc43",
	"8d71c04658ab0455a6f8cfc47575",
	"8d4850b59909a99ed008c28522d8",
	"8d3c4b2e991207004808ca5844dc",
	"5d3c4b2eba4b44",
	"8d44003f9901ec846807ce0d492f",
	"8d44003f9901ec846807ce0d492f",
	"8d3c4b2e58b5010436f7ec50bd5a",
	"8d71c046252c1335cf0820cc7b1d",
	"8d3c4b2e991207004808ca5844dc",
	"8d71c046f82100020049b891fb5a",
	"8d71c0469909f5857808ca17a769",
	"5d4850b5e506cd",
	"8d4850b59909a99ed008c28522d8",
	"8d4850b59909a99ed008c28522d8",
	"5d71c046fdcc43",
	"8d3c4b2e58b5046648b2726aa2cc",
	"5d3c4b2eba4b44",
	"8d3c4b2e58b5046648b28a6f7668",
	"8d3c4b2e991207004808ca5844dc",
	"8d3c4b2ee990b4247138005493a1",
	"5d424444b2b286",
	"8d71c046ea40885cd73c087c186b",
	"8d4850b5e1163e00000000927cfd",
	"8d424444589fe0f6f4e2997ab806",
	"8d71c04658ab045598f926eadb37",
	"8d4850b558b3f04d4f01217563c5",
	"8d424444ea3ca858013c08978b1b",
	"8d4244449909ca84b004c7b235d5",
	"8d4850b59909a99ed004c2cd78d8",
	"5d71c046fdcc43",
	"8d71c0469909f5857804ca5ffd69",
	"5d3c4b2eba4b44",
	"a0001690c4662330aa000028285e",
	"a0001690fff3ff40e004f79b1337",
	"8d3c4b2e2510c237cb0820f55865",
	"a0001690c4662330aa000028285e",
	"8d44003f60b500f8bcd69f099fd2",
	"a0001690c4662330aa000028285e",
	"a000169099ba49356007ff4b01d9",
	"a0001530c0780030a4000029a8b5",
	"8d3c4b2e58b5046648b2b66e1f5a",
	"a0001530c0780030a4000029a8b5",
	"a0001530fff4373ec004ef41961c",
	"a0001530fff4373ec004ef41961c",
	"a00015309bfa4b337fe7ff006381",
	"a0001530fff4373ec004ef41961c",
	"5d3c4b2eba4b44",
	"5d71c046fdcc43",
	"a0001530fff4373ec004ef41961c",
	"5d44003f27cf64",
	"5d71c046fdcc43",
	"8d3c4b2e991207004808ca5844dc",
	"5d3c4b2eba4b44",
	"5d71c046fdcc43",
	"a0001530fff4373ec004ef41961c",
	"5d44003f27cf64",
	"5d71c046fdcc43",
	"5d3c4b2eba4b44",
	"5d424444b2b286",
	"5d424444b2b286",
	"5d424444b2b286",
	"5d68326dfab8eb",
	"8d68326d90b51691d5651f86c375",
	"8d71c04658ab00f3373ef52bde4a",
	"8d424444589fe459389dd2b9d058",
	"8d4244449909ca84b004c7b235d5",
	"8d4244449909ca84b004c7b235d5",
	"8d71c0469909f5857804ca5ffd69",
	"a000163fffd5573d7ffce1f87783",
	"a000163fa55a0f30a07c01482f7a",
	"8d3c4b2e991207004808ca5844dc",
	"a000163fa55a0f30bff401539881",
	"a000163fc4600030aa00003f02ed",
	"a000163fc4600030aa00003f02ed",
}
