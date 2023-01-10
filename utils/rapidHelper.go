package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"shazam_music_query/initializers"
)

func MusicSearch(songTitle string) ([]byte, error) {
	baseURL := "https://shazam.p.rapidapi.com"
	resource := "/search"

	config, _ := initializers.LoadConfig(".")

	// Set up the query parameters
	params := url.Values{}
	params.Add("term", songTitle)

	// Create the URL
	u, _ := url.Parse(baseURL + resource)
	u.RawQuery = params.Encode()

	req, _ := http.NewRequest("GET", u.String(), nil)

	apiKey := config.RapidApiKey

	req.Header.Add("X-RapidAPI-Key", apiKey)
	req.Header.Add("X-RapidAPI-Host", "shazam.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)
	// body, _ := ioutil.ReadAll(res.Body)

	responseData, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return []byte(""), err
	}

	// Check the status code
	if res.StatusCode != http.StatusOK {
		return []byte(""), fmt.Errorf("shazam search failed: %s", res.Status)
	}

	defer res.Body.Close()

	// fmt.Println(string(responseData))
	// fmt.Println(string(body))

	// fmt.Println("responseData:", string(responseData))

	return responseData, nil

	// SAMPLE RESPONSE
	// {
	// 	"tracks":{
	// 	   "hits":[
	// 		  {
	// 			 "track":{
	// 				"layout":"5",
	// 				"type":"MUSIC",
	// 				"key":"556322673",
	// 				"title":"Green Means Go (feat. Rjz & Phyno)",
	// 				"subtitle":"Pappy Kojo",
	// 				"share":{
	// 				   "subject":"Green Means Go (feat. Rjz & Phyno) - Pappy Kojo",
	// 				   "text":"I used Shazam to discover Green Means Go (feat. Rjz & Phyno) by Pappy Kojo.",
	// 				   "href":"https://www.shazam.com/track/556322673/green-means-go-feat-rjz-phyno",
	// 				   "image":"https://is3-ssl.mzstatic.com/image/thumb/Music114/v4/60/2c/ce/602cce68-9d0a-9007-9445-c851254faa2e/0757572917068.png/400x400cc.jpg",
	// 				   "twitter":"I used @Shazam to discover Green Means Go (feat. Rjz & Phyno) by Pappy Kojo.",
	// 				   "html":"https://www.shazam.com/snippets/email-share/556322673?lang=en-US&country=US",
	// 				   "avatar":"https://is1-ssl.mzstatic.com/image/thumb/Music116/v4/40/e3/b2/40e3b2b9-6bea-fbe9-dc3f-94e8330d6f27/pr_source.png/800x800cc.jpg",
	// 				   "snapchat":"https://www.shazam.com/partner/sc/track/556322673"
	// 				},
	// 				"images":{
	// 				   "background":"https://is1-ssl.mzstatic.com/image/thumb/Music116/v4/40/e3/b2/40e3b2b9-6bea-fbe9-dc3f-94e8330d6f27/pr_source.png/800x800cc.jpg",
	// 				   "coverart":"https://is3-ssl.mzstatic.com/image/thumb/Music114/v4/60/2c/ce/602cce68-9d0a-9007-9445-c851254faa2e/0757572917068.png/400x400cc.jpg",
	// 				   "coverarthq":"https://is3-ssl.mzstatic.com/image/thumb/Music114/v4/60/2c/ce/602cce68-9d0a-9007-9445-c851254faa2e/0757572917068.png/400x400cc.jpg",
	// 				   "joecolor":"b:bfc8cdp:080b0cs:282116t:2d3133q:46433b"
	// 				},
	// 				"hub":{
	// 				   "type":"APPLEMUSIC",
	// 				   "image":"https://images.shazam.com/static/icons/hub/ios/v5/applemusic_{scalefactor}.png",
	// 				   "actions":[
	// 					  {
	// 						 "name":"apple",
	// 						 "type":"applemusicplay",
	// 						 "id":"1551986696"
	// 					  },
	// 					  {
	// 						 "name":"apple",
	// 						 "type":"uri",
	// 						 "uri":"https://audio-ssl.itunes.apple.com/itunes-assets/AudioPreview114/v4/fd/02/82/fd028220-5ee6-2dc5-ad55-0b296b25f7d5/mzaf_13966152900860412744.plus.aac.ep.m4a"
	// 					  }
	// 				   ],
	// 				   "options":[
	// 					  {
	// 						 "caption":"OPEN",
	// 						 "actions":[
	// 							{
	// 							   "name":"hub:applemusic:deeplink",
	// 							   "type":"applemusicopen",
	// 							   "uri":"https://music.apple.com/us/album/green-means-go-feat-rjz-phyno/1551986688?i=1551986696&mttnagencyid=s2n&mttnsiteid=125115&mttn3pid=Apple-Shazam&mttnsub1=Shazam_ios&mttnsub2=5348615A-616D-3235-3830-44754D6D5973&itscg=30201&app=music&itsct=Shazam_ios"
	// 							},
	// 							{
	// 							   "name":"hub:applemusic:deeplink",
	// 							   "type":"uri",
	// 							   "uri":"https://music.apple.com/us/album/green-means-go-feat-rjz-phyno/1551986688?i=1551986696&mttnagencyid=s2n&mttnsiteid=125115&mttn3pid=Apple-Shazam&mttnsub1=Shazam_ios&mttnsub2=5348615A-616D-3235-3830-44754D6D5973&itscg=30201&app=music&itsct=Shazam_ios"
	// 							}
	// 						 ],
	// 						 "beacondata":{
	// 							"type":"open",
	// 							"providername":"applemusic"
	// 						 },
	// 						 "image":"https://images.shazam.com/static/icons/hub/ios/v5/overflow-open-option_{scalefactor}.png",
	// 						 "type":"open",
	// 						 "listcaption":"Open in Apple Music",
	// 						 "overflowimage":"https://images.shazam.com/static/icons/hub/ios/v5/applemusic-overflow_{scalefactor}.png",
	// 						 "colouroverflowimage":false,
	// 						 "providername":"applemusic"
	// 					  },
	// 					  {
	// 						 "caption":"BUY",
	// 						 "actions":[
	// 							{
	// 							   "type":"uri",
	// 							   "uri":"https://itunes.apple.com/us/album/green-means-go-feat-rjz-phyno/1551986688?i=1551986696&mttnagencyid=s2n&mttnsiteid=125115&mttn3pid=Apple-Shazam&mttnsub1=Shazam_ios&mttnsub2=5348615A-616D-3235-3830-44754D6D5973&itscg=30201&app=itunes&itsct=Shazam_ios"
	// 							}
	// 						 ],
	// 						 "beacondata":{
	// 							"type":"buy",
	// 							"providername":"itunes"
	// 						 },
	// 						 "image":"https://images.shazam.com/static/icons/hub/ios/v5/itunes-overflow-buy_{scalefactor}.png",
	// 						 "type":"buy",
	// 						 "listcaption":"Buy on iTunes",
	// 						 "overflowimage":"https://images.shazam.com/static/icons/hub/ios/v5/itunes-overflow-buy_{scalefactor}.png",
	// 						 "colouroverflowimage":false,
	// 						 "providername":"itunes"
	// 					  }
	// 				   ],
	// 				   "providers":[
	// 					  {
	// 						 "caption":"Open in Spotify",
	// 						 "images":{
	// 							"overflow":"https://images.shazam.com/static/icons/hub/ios/v5/spotify-overflow_{scalefactor}.png",
	// 							"default":"https://images.shazam.com/static/icons/hub/ios/v5/spotify_{scalefactor}.png"
	// 						 },
	// 						 "actions":[
	// 							{
	// 							   "name":"hub:spotify:searchdeeplink",
	// 							   "type":"uri",
	// 							   "uri":"spotify:search:Green%20Means%20Go%20%28feat.%20Rjz%20%26%20Phyno%29%20Pappy%20Kojo"
	// 							}
	// 						 ],
	// 						 "type":"SPOTIFY"
	// 					  },
	// 					  {
	// 						 "caption":"Open in Deezer",
	// 						 "images":{
	// 							"overflow":"https://images.shazam.com/static/icons/hub/ios/v5/deezer-overflow_{scalefactor}.png",
	// 							"default":"https://images.shazam.com/static/icons/hub/ios/v5/deezer_{scalefactor}.png"
	// 						 },
	// 						 "actions":[
	// 							{
	// 							   "name":"hub:deezer:searchdeeplink",
	// 							   "type":"uri",
	// 							   "uri":"deezer-query://www.deezer.com/play?query=%7Btrack%3A%27Green+Means+Go+%28feat.+Rjz++Phyno%29%27%20artist%3A%27Pappy+Kojo%27%7D"
	// 							}
	// 						 ],
	// 						 "type":"DEEZER"
	// 					  }
	// 				   ],
	// 				   "explicit":true,
	// 				   "displayname":"APPLE MUSIC"
	// 				},
	// 				"artists":[
	// 				   {
	// 					  "id":"42",
	// 					  "adamid":"943085329"
	// 				   }
	// 				],
	// 				"url":"https://www.shazam.com/track/556322673/green-means-go-feat-rjz-phyno"
	// 			 }
	// 		  },
	// 		  {
	// 			 "track":{
	// 				"layout":"5",
	// 				"type":"MUSIC",
	// 				"key":"471317212",
	// 				"title":"Green Means Go (feat. RJZ)",
	// 				"subtitle":"Pappy Kojo",
	// 				"share":{
	// 				   "subject":"Green Means Go (feat. RJZ) - Pappy Kojo",
	// 				   "text":"I used Shazam to discover Green Means Go (feat. RJZ) by Pappy Kojo.",
	// 				   "href":"https://www.shazam.com/track/471317212/green-means-go-feat-rjz",
	// 				   "image":"https://is4-ssl.mzstatic.com/image/thumb/Music123/v4/7a/82/be/7a82be75-3a9f-c532-a863-9111289616d0/0757572845279.png/400x400cc.jpg",
	// 				   "twitter":"I used @Shazam to discover Green Means Go (feat. RJZ) by Pappy Kojo.",
	// 				   "html":"https://www.shazam.com/snippets/email-share/471317212?lang=en-US&country=US",
	// 				   "avatar":"https://is1-ssl.mzstatic.com/image/thumb/Music116/v4/40/e3/b2/40e3b2b9-6bea-fbe9-dc3f-94e8330d6f27/pr_source.png/800x800cc.jpg",
	// 				   "snapchat":"https://www.shazam.com/partner/sc/track/471317212"
	// 				},
	// 				"images":{
	// 				   "background":"https://is1-ssl.mzstatic.com/image/thumb/Music116/v4/40/e3/b2/40e3b2b9-6bea-fbe9-dc3f-94e8330d6f27/pr_source.png/800x800cc.jpg",
	// 				   "coverart":"https://is4-ssl.mzstatic.com/image/thumb/Music123/v4/7a/82/be/7a82be75-3a9f-c532-a863-9111289616d0/0757572845279.png/400x400cc.jpg",
	// 				   "coverarthq":"https://is4-ssl.mzstatic.com/image/thumb/Music123/v4/7a/82/be/7a82be75-3a9f-c532-a863-9111289616d0/0757572845279.png/400x400cc.jpg",
	// 				   "joecolor":"b:4f9d73p:080300s:151c0bt:162217q:213620"
	// 				},
	// 				"hub":{
	// 				   "type":"APPLEMUSIC",
	// 				   "image":"https://images.shazam.com/static/icons/hub/ios/v5/applemusic_{scalefactor}.png",
	// 				   "actions":[
	// 					  {
	// 						 "name":"apple",
	// 						 "type":"applemusicplay",
	// 						 "id":"1471965314"
	// 					  },
	// 					  {
	// 						 "name":"apple",
	// 						 "type":"uri",
	// 						 "uri":"https://audio-ssl.itunes.apple.com/itunes-assets/AudioPreview115/v4/dd/49/cd/dd49cd63-0143-0dc1-a0ba-44c9ee642ab1/mzaf_8279158982791141821.plus.aac.p.m4a"
	// 					  }
	// 				   ],
	// 				   "options":[
	// 					  {
	// 						 "caption":"OPEN",
	// 						 "actions":[
	// 							{
	// 							   "name":"hub:applemusic:deeplink",
	// 							   "type":"applemusicopen",
	// 							   "uri":"https://music.apple.com/us/album/green-means-go-feat-rjz/1471965312?i=1471965314&mttnagencyid=s2n&mttnsiteid=125115&mttn3pid=Apple-Shazam&mttnsub1=Shazam_ios&mttnsub2=5348615A-616D-3235-3830-44754D6D5973&itscg=30201&app=music&itsct=Shazam_ios"
	// 							},
	// 							{
	// 							   "name":"hub:applemusic:deeplink",
	// 							   "type":"uri",
	// 							   "uri":"https://music.apple.com/us/album/green-means-go-feat-rjz/1471965312?i=1471965314&mttnagencyid=s2n&mttnsiteid=125115&mttn3pid=Apple-Shazam&mttnsub1=Shazam_ios&mttnsub2=5348615A-616D-3235-3830-44754D6D5973&itscg=30201&app=music&itsct=Shazam_ios"
	// 							}
	// 						 ],
	// 						 "beacondata":{
	// 							"type":"open",
	// 							"providername":"applemusic"
	// 						 },
	// 						 "image":"https://images.shazam.com/static/icons/hub/ios/v5/overflow-open-option_{scalefactor}.png",
	// 						 "type":"open",
	// 						 "listcaption":"Open in Apple Music",
	// 						 "overflowimage":"https://images.shazam.com/static/icons/hub/ios/v5/applemusic-overflow_{scalefactor}.png",
	// 						 "colouroverflowimage":false,
	// 						 "providername":"applemusic"
	// 					  },
	// 					  {
	// 						 "caption":"BUY",
	// 						 "actions":[
	// 							{
	// 							   "type":"uri",
	// 							   "uri":"https://itunes.apple.com/us/album/green-means-go-feat-rjz/1471965312?i=1471965314&mttnagencyid=s2n&mttnsiteid=125115&mttn3pid=Apple-Shazam&mttnsub1=Shazam_ios&mttnsub2=5348615A-616D-3235-3830-44754D6D5973&itscg=30201&app=itunes&itsct=Shazam_ios"
	// 							}
	// 						 ],
	// 						 "beacondata":{
	// 							"type":"buy",
	// 							"providername":"itunes"
	// 						 },
	// 						 "image":"https://images.shazam.com/static/icons/hub/ios/v5/itunes-overflow-buy_{scalefactor}.png",
	// 						 "type":"buy",
	// 						 "listcaption":"Buy on iTunes",
	// 						 "overflowimage":"https://images.shazam.com/static/icons/hub/ios/v5/itunes-overflow-buy_{scalefactor}.png",
	// 						 "colouroverflowimage":false,
	// 						 "providername":"itunes"
	// 					  }
	// 				   ],
	// 				   "providers":[
	// 					  {
	// 						 "caption":"Open in Spotify",
	// 						 "images":{
	// 							"overflow":"https://images.shazam.com/static/icons/hub/ios/v5/spotify-overflow_{scalefactor}.png",
	// 							"default":"https://images.shazam.com/static/icons/hub/ios/v5/spotify_{scalefactor}.png"
	// 						 },
	// 						 "actions":[
	// 							{
	// 							   "name":"hub:spotify:searchdeeplink",
	// 							   "type":"uri",
	// 							   "uri":"spotify:search:Green%20Means%20Go%20%28feat.%20RJZ%29%20Pappy%20Kojo"
	// 							}
	// 						 ],
	// 						 "type":"SPOTIFY"
	// 					  },
	// 					  {
	// 						 "caption":"Open in Deezer",
	// 						 "images":{
	// 							"overflow":"https://images.shazam.com/static/icons/hub/ios/v5/deezer-overflow_{scalefactor}.png",
	// 							"default":"https://images.shazam.com/static/icons/hub/ios/v5/deezer_{scalefactor}.png"
	// 						 },
	// 						 "actions":[
	// 							{
	// 							   "name":"hub:deezer:searchdeeplink",
	// 							   "type":"uri",
	// 							   "uri":"deezer-query://www.deezer.com/play?query=%7Btrack%3A%27Green+Means+Go+%28feat.+RJZ%29%27%20artist%3A%27Pappy+Kojo%27%7D"
	// 							}
	// 						 ],
	// 						 "type":"DEEZER"
	// 					  }
	// 				   ],
	// 				   "explicit":false,
	// 				   "displayname":"APPLE MUSIC"
	// 				},
	// 				"artists":[
	// 				   {
	// 					  "id":"42",
	// 					  "adamid":"943085329"
	// 				   }
	// 				],
	// 				"url":"https://www.shazam.com/track/471317212/green-means-go-feat-rjz"
	// 			 }
	// 		  },
	// 		  {
	// 			 "track":{
	// 				"layout":"5",
	// 				"type":"MUSIC",
	// 				"key":"44494084",
	// 				"title":"Green Light",
	// 				"subtitle":"Beyoncé",
	// 				"share":{
	// 				   "subject":"Green Light - Beyoncé",
	// 				   "text":"I used Shazam to discover Green Light by Beyoncé.",
	// 				   "href":"https://www.shazam.com/track/44494084/green-light",
	// 				   "image":"https://is4-ssl.mzstatic.com/image/thumb/Music125/v4/47/b2/b2/47b2b26b-5abd-0953-6688-7dfff80508ec/mzi.sozuruch.jpg/400x400cc.jpg",
	// 				   "twitter":"I used @Shazam to discover Green Light by Beyoncé.",
	// 				   "html":"https://www.shazam.com/snippets/email-share/44494084?lang=en-US&country=US",
	// 				   "avatar":"https://is5-ssl.mzstatic.com/image/thumb/Features122/v4/e2/10/a7/e210a754-3409-4e42-8fd9-413c1289cbb9/mza_3319038547447377908.png/800x800cc.jpg",
	// 				   "snapchat":"https://www.shazam.com/partner/sc/track/44494084"
	// 				},
	// 				"images":{
	// 				   "background":"https://is5-ssl.mzstatic.com/image/thumb/Features122/v4/e2/10/a7/e210a754-3409-4e42-8fd9-413c1289cbb9/mza_3319038547447377908.png/800x800cc.jpg",
	// 				   "coverart":"https://is4-ssl.mzstatic.com/image/thumb/Music125/v4/47/b2/b2/47b2b26b-5abd-0953-6688-7dfff80508ec/mzi.sozuruch.jpg/400x400cc.jpg",
	// 				   "coverarthq":"https://is4-ssl.mzstatic.com/image/thumb/Music125/v4/47/b2/b2/47b2b26b-5abd-0953-6688-7dfff80508ec/mzi.sozuruch.jpg/400x400cc.jpg",
	// 				   "joecolor":"b:b1dbe5p:040405s:492616t:272f32q:5e4a3f"
	// 				},
	// 				"hub":{
	// 				   "type":"APPLEMUSIC",
	// 				   "image":"https://images.shazam.com/static/icons/hub/ios/v5/applemusic_{scalefactor}.png",
	// 				   "actions":[
	// 					  {
	// 						 "name":"apple",
	// 						 "type":"applemusicplay",
	// 						 "id":"261707099"
	// 					  },
	// 					  {
	// 						 "name":"apple",
	// 						 "type":"uri",
	// 						 "uri":"https://audio-ssl.itunes.apple.com/itunes-assets/AudioPreview115/v4/51/04/22/51042279-f9b1-4a54-28b9-082bbaa221ca/mzaf_10372343078745401267.plus.aac.ep.m4a"
	// 					  }
	// 				   ],
	// 				   "options":[
	// 					  {
	// 						 "caption":"OPEN",
	// 						 "actions":[
	// 							{
	// 							   "name":"hub:applemusic:deeplink",
	// 							   "type":"applemusicopen",
	// 							   "uri":"https://music.apple.com/us/album/green-light/261707051?i=261707099&mttnagencyid=s2n&mttnsiteid=125115&mttn3pid=Apple-Shazam&mttnsub1=Shazam_ios&mttnsub2=5348615A-616D-3235-3830-44754D6D5973&itscg=30201&app=music&itsct=Shazam_ios"
	// 							},
	// 							{
	// 							   "name":"hub:applemusic:deeplink",
	// 							   "type":"uri",
	// 							   "uri":"https://music.apple.com/us/album/green-light/261707051?i=261707099&mttnagencyid=s2n&mttnsiteid=125115&mttn3pid=Apple-Shazam&mttnsub1=Shazam_ios&mttnsub2=5348615A-616D-3235-3830-44754D6D5973&itscg=30201&app=music&itsct=Shazam_ios"
	// 							}
	// 						 ],
	// 						 "beacondata":{
	// 							"type":"open",
	// 							"providername":"applemusic"
	// 						 },
	// 						 "image":"https://images.shazam.com/static/icons/hub/ios/v5/overflow-open-option_{scalefactor}.png",
	// 						 "type":"open",
	// 						 "listcaption":"Open in Apple Music",
	// 						 "overflowimage":"https://images.shazam.com/static/icons/hub/ios/v5/applemusic-overflow_{scalefactor}.png",
	// 						 "colouroverflowimage":false,
	// 						 "providername":"applemusic"
	// 					  },
	// 					  {
	// 						 "caption":"BUY",
	// 						 "actions":[
	// 							{
	// 							   "type":"uri",
	// 							   "uri":"https://itunes.apple.com/us/album/green-light/261707051?i=261707099&mttnagencyid=s2n&mttnsiteid=125115&mttn3pid=Apple-Shazam&mttnsub1=Shazam_ios&mttnsub2=5348615A-616D-3235-3830-44754D6D5973&itscg=30201&app=itunes&itsct=Shazam_ios"
	// 							}
	// 						 ],
	// 						 "beacondata":{
	// 							"type":"buy",
	// 							"providername":"itunes"
	// 						 },
	// 						 "image":"https://images.shazam.com/static/icons/hub/ios/v5/itunes-overflow-buy_{scalefactor}.png",
	// 						 "type":"buy",
	// 						 "listcaption":"Buy on iTunes",
	// 						 "overflowimage":"https://images.shazam.com/static/icons/hub/ios/v5/itunes-overflow-buy_{scalefactor}.png",
	// 						 "colouroverflowimage":false,
	// 						 "providername":"itunes"
	// 					  }
	// 				   ],
	// 				   "providers":[
	// 					  {
	// 						 "caption":"Open in Spotify",
	// 						 "images":{
	// 							"overflow":"https://images.shazam.com/static/icons/hub/ios/v5/spotify-overflow_{scalefactor}.png",
	// 							"default":"https://images.shazam.com/static/icons/hub/ios/v5/spotify_{scalefactor}.png"
	// 						 },
	// 						 "actions":[
	// 							{
	// 							   "name":"hub:spotify:searchdeeplink",
	// 							   "type":"uri",
	// 							   "uri":"spotify:search:Green%20Light%20Beyonc%C3%A9"
	// 							}
	// 						 ],
	// 						 "type":"SPOTIFY"
	// 					  },
	// 					  {
	// 						 "caption":"Open in Deezer",
	// 						 "images":{
	// 							"overflow":"https://images.shazam.com/static/icons/hub/ios/v5/deezer-overflow_{scalefactor}.png",
	// 							"default":"https://images.shazam.com/static/icons/hub/ios/v5/deezer_{scalefactor}.png"
	// 						 },
	// 						 "actions":[
	// 							{
	// 							   "name":"hub:deezer:searchdeeplink",
	// 							   "type":"uri",
	// 							   "uri":"deezer-query://www.deezer.com/play?query=%7Btrack%3A%27Green+Light%27%20artist%3A%27Beyonc%C3%A9%27%7D"
	// 							}
	// 						 ],
	// 						 "type":"DEEZER"
	// 					  }
	// 				   ],
	// 				   "explicit":false,
	// 				   "displayname":"APPLE MUSIC"
	// 				},
	// 				"artists":[
	// 				   {
	// 					  "id":"42",
	// 					  "adamid":"1419227"
	// 				   }
	// 				],
	// 				"url":"https://www.shazam.com/track/44494084/green-light"
	// 			 },
	// 			 "snippet":"You holding up traffic, green means go"
	// 		  },
	// 		  {
	// 			 "track":{
	// 				"layout":"5",
	// 				"type":"MUSIC",
	// 				"key":"565051905",
	// 				"title":"Kumala Vista",
	// 				"subtitle":"Green Means Go",
	// 				"share":{
	// 				   "subject":"Kumala Vista - Green Means Go",
	// 				   "text":"I used Shazam to discover Kumala Vista by Green Means Go.",
	// 				   "href":"https://www.shazam.com/track/565051905/kumala-vista",
	// 				   "image":"https://is2-ssl.mzstatic.com/image/thumb/Music114/v4/d5/fb/d4/d5fbd444-ce22-524d-4c85-ef0fd7a5c931/672985839717.png/400x400cc.jpg",
	// 				   "twitter":"I used @Shazam to discover Kumala Vista by Green Means Go.",
	// 				   "html":"https://www.shazam.com/snippets/email-share/565051905?lang=en-US&country=US",
	// 				   "snapchat":"https://www.shazam.com/partner/sc/track/565051905"
	// 				},
	// 				"images":{
	// 				   "background":"https://is2-ssl.mzstatic.com/image/thumb/Music114/v4/d5/fb/d4/d5fbd444-ce22-524d-4c85-ef0fd7a5c931/672985839717.png/400x400cc.jpg",
	// 				   "coverart":"https://is2-ssl.mzstatic.com/image/thumb/Music114/v4/d5/fb/d4/d5fbd444-ce22-524d-4c85-ef0fd7a5c931/672985839717.png/400x400cc.jpg",
	// 				   "coverarthq":"https://is2-ssl.mzstatic.com/image/thumb/Music114/v4/d5/fb/d4/d5fbd444-ce22-524d-4c85-ef0fd7a5c931/672985839717.png/400x400cc.jpg",
	// 				   "joecolor":"b:06bd87p:000101s:051e20t:01261cq:053d34"
	// 				},
	// 				"hub":{
	// 				   "type":"APPLEMUSIC",
	// 				   "image":"https://images.shazam.com/static/icons/hub/ios/v5/applemusic_{scalefactor}.png",
	// 				   "actions":[
	// 					  {
	// 						 "name":"apple",
	// 						 "type":"applemusicplay",
	// 						 "id":"1561802992"
	// 					  },
	// 					  {
	// 						 "name":"apple",
	// 						 "type":"uri",
	// 						 "uri":"https://audio-ssl.itunes.apple.com/itunes-assets/AudioPreview114/v4/c1/3c/3e/c13c3e7e-f435-ba23-5282-11b7cfbf2f32/mzaf_5617803953389401545.plus.aac.p.m4a"
	// 					  }
	// 				   ],
	// 				   "options":[
	// 					  {
	// 						 "caption":"OPEN",
	// 						 "actions":[
	// 							{
	// 							   "name":"hub:applemusic:deeplink",
	// 							   "type":"applemusicopen",
	// 							   "uri":"https://music.apple.com/us/album/kumala-vista/1561802978?i=1561802992&mttnagencyid=s2n&mttnsiteid=125115&mttn3pid=Apple-Shazam&mttnsub1=Shazam_ios&mttnsub2=5348615A-616D-3235-3830-44754D6D5973&itscg=30201&app=music&itsct=Shazam_ios"
	// 							},
	// 							{
	// 							   "name":"hub:applemusic:deeplink",
	// 							   "type":"uri",
	// 							   "uri":"https://music.apple.com/us/album/kumala-vista/1561802978?i=1561802992&mttnagencyid=s2n&mttnsiteid=125115&mttn3pid=Apple-Shazam&mttnsub1=Shazam_ios&mttnsub2=5348615A-616D-3235-3830-44754D6D5973&itscg=30201&app=music&itsct=Shazam_ios"
	// 							}
	// 						 ],
	// 						 "beacondata":{
	// 							"type":"open",
	// 							"providername":"applemusic"
	// 						 },
	// 						 "image":"https://images.shazam.com/static/icons/hub/ios/v5/overflow-open-option_{scalefactor}.png",
	// 						 "type":"open",
	// 						 "listcaption":"Open in Apple Music",
	// 						 "overflowimage":"https://images.shazam.com/static/icons/hub/ios/v5/applemusic-overflow_{scalefactor}.png",
	// 						 "colouroverflowimage":false,
	// 						 "providername":"applemusic"
	// 					  },
	// 					  {
	// 						 "caption":"BUY",
	// 						 "actions":[
	// 							{
	// 							   "type":"uri",
	// 							   "uri":"https://itunes.apple.com/us/album/kumala-vista/1561802978?i=1561802992&mttnagencyid=s2n&mttnsiteid=125115&mttn3pid=Apple-Shazam&mttnsub1=Shazam_ios&mttnsub2=5348615A-616D-3235-3830-44754D6D5973&itscg=30201&app=itunes&itsct=Shazam_ios"
	// 							}
	// 						 ],
	// 						 "beacondata":{
	// 							"type":"buy",
	// 							"providername":"itunes"
	// 						 },
	// 						 "image":"https://images.shazam.com/static/icons/hub/ios/v5/itunes-overflow-buy_{scalefactor}.png",
	// 						 "type":"buy",
	// 						 "listcaption":"Buy on iTunes",
	// 						 "overflowimage":"https://images.shazam.com/static/icons/hub/ios/v5/itunes-overflow-buy_{scalefactor}.png",
	// 						 "colouroverflowimage":false,
	// 						 "providername":"itunes"
	// 					  }
	// 				   ],
	// 				   "providers":[
	// 					  {
	// 						 "caption":"Open in Spotify",
	// 						 "images":{
	// 							"overflow":"https://images.shazam.com/static/icons/hub/ios/v5/spotify-overflow_{scalefactor}.png",
	// 							"default":"https://images.shazam.com/static/icons/hub/ios/v5/spotify_{scalefactor}.png"
	// 						 },
	// 						 "actions":[
	// 							{
	// 							   "name":"hub:spotify:searchdeeplink",
	// 							   "type":"uri",
	// 							   "uri":"spotify:search:Kumala%20Vista%20Green%20Means%20Go"
	// 							}
	// 						 ],
	// 						 "type":"SPOTIFY"
	// 					  },
	// 					  {
	// 						 "caption":"Open in Deezer",
	// 						 "images":{
	// 							"overflow":"https://images.shazam.com/static/icons/hub/ios/v5/deezer-overflow_{scalefactor}.png",
	// 							"default":"https://images.shazam.com/static/icons/hub/ios/v5/deezer_{scalefactor}.png"
	// 						 },
	// 						 "actions":[
	// 							{
	// 							   "name":"hub:deezer:searchdeeplink",
	// 							   "type":"uri",
	// 							   "uri":"deezer-query://www.deezer.com/play?query=%7Btrack%3A%27Kumala+Vista%27%20artist%3A%27Green+Means+Go%27%7D"
	// 							}
	// 						 ],
	// 						 "type":"DEEZER"
	// 					  }
	// 				   ],
	// 				   "explicit":false,
	// 				   "displayname":"APPLE MUSIC"
	// 				},
	// 				"artists":[
	// 				   {
	// 					  "id":"42",
	// 					  "adamid":"311431960"
	// 				   }
	// 				],
	// 				"url":"https://www.shazam.com/track/565051905/kumala-vista"
	// 			 }
	// 		  },
	// 		  {
	// 			 "track":{
	// 				"layout":"5",
	// 				"type":"MUSIC",
	// 				"key":"565051907",
	// 				"title":"Deep in the Jungle",
	// 				"subtitle":"Green Means Go",
	// 				"share":{
	// 				   "subject":"Deep in the Jungle - Green Means Go",
	// 				   "text":"I used Shazam to discover Deep in the Jungle by Green Means Go.",
	// 				   "href":"https://www.shazam.com/track/565051907/deep-in-the-jungle",
	// 				   "image":"https://is2-ssl.mzstatic.com/image/thumb/Music114/v4/d5/fb/d4/d5fbd444-ce22-524d-4c85-ef0fd7a5c931/672985839717.png/400x400cc.jpg",
	// 				   "twitter":"I used @Shazam to discover Deep in the Jungle by Green Means Go.",
	// 				   "html":"https://www.shazam.com/snippets/email-share/565051907?lang=en-US&country=US",
	// 				   "snapchat":"https://www.shazam.com/partner/sc/track/565051907"
	// 				},
	// 				"images":{
	// 				   "background":"https://is2-ssl.mzstatic.com/image/thumb/Music114/v4/d5/fb/d4/d5fbd444-ce22-524d-4c85-ef0fd7a5c931/672985839717.png/400x400cc.jpg",
	// 				   "coverart":"https://is2-ssl.mzstatic.com/image/thumb/Music114/v4/d5/fb/d4/d5fbd444-ce22-524d-4c85-ef0fd7a5c931/672985839717.png/400x400cc.jpg",
	// 				   "coverarthq":"https://is2-ssl.mzstatic.com/image/thumb/Music114/v4/d5/fb/d4/d5fbd444-ce22-524d-4c85-ef0fd7a5c931/672985839717.png/400x400cc.jpg",
	// 				   "joecolor":"b:06bd87p:000101s:051e20t:01261cq:053d34"
	// 				},
	// 				"hub":{
	// 				   "type":"APPLEMUSIC",
	// 				   "image":"https://images.shazam.com/static/icons/hub/ios/v5/applemusic_{scalefactor}.png",
	// 				   "actions":[
	// 					  {
	// 						 "name":"apple",
	// 						 "type":"applemusicplay",
	// 						 "id":"1561802982"
	// 					  },
	// 					  {
	// 						 "name":"apple",
	// 						 "type":"uri",
	// 						 "uri":"https://audio-ssl.itunes.apple.com/itunes-assets/AudioPreview114/v4/2e/42/a3/2e42a35f-0904-ef6b-71ee-2f7ab9a8a675/mzaf_13650749200726605957.plus.aac.p.m4a"
	// 					  }
	// 				   ],
	// 				   "options":[
	// 					  {
	// 						 "caption":"OPEN",
	// 						 "actions":[
	// 							{
	// 							   "name":"hub:applemusic:deeplink",
	// 							   "type":"applemusicopen",
	// 							   "uri":"https://music.apple.com/us/album/deep-in-the-jungle/1561802978?i=1561802982&mttnagencyid=s2n&mttnsiteid=125115&mttn3pid=Apple-Shazam&mttnsub1=Shazam_ios&mttnsub2=5348615A-616D-3235-3830-44754D6D5973&itscg=30201&app=music&itsct=Shazam_ios"
	// 							},
	// 							{
	// 							   "name":"hub:applemusic:deeplink",
	// 							   "type":"uri",
	// 							   "uri":"https://music.apple.com/us/album/deep-in-the-jungle/1561802978?i=1561802982&mttnagencyid=s2n&mttnsiteid=125115&mttn3pid=Apple-Shazam&mttnsub1=Shazam_ios&mttnsub2=5348615A-616D-3235-3830-44754D6D5973&itscg=30201&app=music&itsct=Shazam_ios"
	// 							}
	// 						 ],
	// 						 "beacondata":{
	// 							"type":"open",
	// 							"providername":"applemusic"
	// 						 },
	// 						 "image":"https://images.shazam.com/static/icons/hub/ios/v5/overflow-open-option_{scalefactor}.png",
	// 						 "type":"open",
	// 						 "listcaption":"Open in Apple Music",
	// 						 "overflowimage":"https://images.shazam.com/static/icons/hub/ios/v5/applemusic-overflow_{scalefactor}.png",
	// 						 "colouroverflowimage":false,
	// 						 "providername":"applemusic"
	// 					  },
	// 					  {
	// 						 "caption":"BUY",
	// 						 "actions":[
	// 							{
	// 							   "type":"uri",
	// 							   "uri":"https://itunes.apple.com/us/album/deep-in-the-jungle/1561802978?i=1561802982&mttnagencyid=s2n&mttnsiteid=125115&mttn3pid=Apple-Shazam&mttnsub1=Shazam_ios&mttnsub2=5348615A-616D-3235-3830-44754D6D5973&itscg=30201&app=itunes&itsct=Shazam_ios"
	// 							}
	// 						 ],
	// 						 "beacondata":{
	// 							"type":"buy",
	// 							"providername":"itunes"
	// 						 },
	// 						 "image":"https://images.shazam.com/static/icons/hub/ios/v5/itunes-overflow-buy_{scalefactor}.png",
	// 						 "type":"buy",
	// 						 "listcaption":"Buy on iTunes",
	// 						 "overflowimage":"https://images.shazam.com/static/icons/hub/ios/v5/itunes-overflow-buy_{scalefactor}.png",
	// 						 "colouroverflowimage":false,
	// 						 "providername":"itunes"
	// 					  }
	// 				   ],
	// 				   "providers":[
	// 					  {
	// 						 "caption":"Open in Spotify",
	// 						 "images":{
	// 							"overflow":"https://images.shazam.com/static/icons/hub/ios/v5/spotify-overflow_{scalefactor}.png",
	// 							"default":"https://images.shazam.com/static/icons/hub/ios/v5/spotify_{scalefactor}.png"
	// 						 },
	// 						 "actions":[
	// 							{
	// 							   "name":"hub:spotify:searchdeeplink",
	// 							   "type":"uri",
	// 							   "uri":"spotify:search:Deep%20in%20the%20Jungle%20Green%20Means%20Go"
	// 							}
	// 						 ],
	// 						 "type":"SPOTIFY"
	// 					  },
	// 					  {
	// 						 "caption":"Open in Deezer",
	// 						 "images":{
	// 							"overflow":"https://images.shazam.com/static/icons/hub/ios/v5/deezer-overflow_{scalefactor}.png",
	// 							"default":"https://images.shazam.com/static/icons/hub/ios/v5/deezer_{scalefactor}.png"
	// 						 },
	// 						 "actions":[
	// 							{
	// 							   "name":"hub:deezer:searchdeeplink",
	// 							   "type":"uri",
	// 							   "uri":"deezer-query://www.deezer.com/play?query=%7Btrack%3A%27Deep+in+the+Jungle%27%20artist%3A%27Green+Means+Go%27%7D"
	// 							}
	// 						 ],
	// 						 "type":"DEEZER"
	// 					  }
	// 				   ],
	// 				   "explicit":false,
	// 				   "displayname":"APPLE MUSIC"
	// 				},
	// 				"artists":[
	// 				   {
	// 					  "id":"42",
	// 					  "adamid":"311431960"
	// 				   }
	// 				],
	// 				"url":"https://www.shazam.com/track/565051907/deep-in-the-jungle"
	// 			 }
	// 		  }
	// 	   ]
	// 	},
	// 	"artists":{
	// 	   "hits":[
	// 		  {
	// 			 "artist":{
	// 				"avatar":"https://is4-ssl.mzstatic.com/image/thumb/Music116/v4/40/e3/b2/40e3b2b9-6bea-fbe9-dc3f-94e8330d6f27/pr_source.png/800x800bb.jpg",
	// 				"name":"Pappy Kojo",
	// 				"verified":false,
	// 				"weburl":"https://music.apple.com/us/artist/pappy-kojo/943085329",
	// 				"adamid":"943085329"
	// 			 }
	// 		  },
	// 		  {
	// 			 "artist":{
	// 				"avatar":"https://is2-ssl.mzstatic.com/image/thumb/AMCArtistImages122/v4/b0/d0/ea/b0d0eaf8-654b-d3a1-73f9-3274d857bc89/906c3c9f-dbf6-45f5-b8ae-c8cefe9612ae_file_cropped.png/800x800bb.jpg",
	// 				"name":"Phyno",
	// 				"verified":false,
	// 				"weburl":"https://music.apple.com/us/artist/phyno/514017875",
	// 				"adamid":"514017875"
	// 			 }
	// 		  },
	// 		  {
	// 			 "artist":{
	// 				"avatar":"https://is1-ssl.mzstatic.com/image/thumb/Music112/v4/a1/3f/c0/a13fc08c-a538-72fa-fa3a-1c5299d78cd8/pr_source.png/800x800bb.jpg",
	// 				"name":"Rjz",
	// 				"verified":false,
	// 				"weburl":"https://music.apple.com/us/artist/rjz/1215336127",
	// 				"adamid":"1215336127"
	// 			 }
	// 		  },
	// 		  {
	// 			 "artist":{
	// 				"avatar":"https://is3-ssl.mzstatic.com/image/thumb/Music114/v4/d5/fb/d4/d5fbd444-ce22-524d-4c85-ef0fd7a5c931/672985839717.png/800x800ac.jpg",
	// 				"name":"Green Means Go",
	// 				"verified":false,
	// 				"weburl":"https://music.apple.com/us/artist/green-means-go/311431960",
	// 				"adamid":"311431960"
	// 			 }
	// 		  },
	// 		  {
	// 			 "artist":{
	// 				"avatar":"https://is3-ssl.mzstatic.com/image/thumb/Features122/v4/e2/10/a7/e210a754-3409-4e42-8fd9-413c1289cbb9/mza_3319038547447377908.png/800x800bb.jpg",
	// 				"name":"Beyoncé",
	// 				"verified":false,
	// 				"weburl":"https://music.apple.com/us/artist/beyonc%C3%A9/1419227",
	// 				"adamid":"1419227"
	// 			 }
	// 		  }
	// 	   ]
	// 	}
	//  }

}

func GetTotalStream(musicKey string) ([]byte, error) {
	baseURL := "https://shazam.p.rapidapi.com"
	resource := "/songs/get-count"

	params := url.Values{}
	params.Add("key", musicKey)

	config, _ := initializers.LoadConfig(".")

	// Create the URL
	u, _ := url.Parse(baseURL + resource)
	u.RawQuery = params.Encode()

	req, _ := http.NewRequest("GET", u.String(), nil)

	apiKey := config.RapidApiKey

	req.Header.Add("X-RapidAPI-Key", apiKey)
	req.Header.Add("X-RapidAPI-Host", "shazam.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	responseData, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return []byte(""), err
	}

	defer res.Body.Close()

	return responseData, nil
}
