// This file is generated by running `make generateAPI`
package generated

type Teams struct {
	Teams []struct {
		Abbrev              string   `json:"abbrev"`
		DivisionID          int64    `json:"divisionId"`
		ID                  int64    `json:"id"`
		IsActive            bool     `json:"isActive"`
		Location            string   `json:"location"`
		Logo                string   `json:"logo"`
		LogoType            string   `json:"logoType"`
		Nickname            string   `json:"nickname"`
		Owners              []string `json:"owners"`
		PlayoffSeed         int64    `json:"playoffSeed"`
		Points              float64  `json:"points"`
		PointsAdjusted      float64  `json:"pointsAdjusted"`
		PointsDelta         int64    `json:"pointsDelta"`
		PrimaryOwner        string   `json:"primaryOwner"`
		RankCalculatedFinal int64    `json:"rankCalculatedFinal"`
		RankFinal           int64    `json:"rankFinal"`
		Record              struct {
			Away struct {
				GamesBack     int64  `json:"gamesBack"`
				Losses        int64  `json:"losses"`
				Percentage    int64  `json:"percentage"`
				PointsAgainst int64  `json:"pointsAgainst"`
				PointsFor     int64  `json:"pointsFor"`
				StreakLength  int64  `json:"streakLength"`
				StreakType    string `json:"streakType"`
				Ties          int64  `json:"ties"`
				Wins          int64  `json:"wins"`
			} `json:"away"`
			Division struct {
				GamesBack     int64   `json:"gamesBack"`
				Losses        int64   `json:"losses"`
				Percentage    float64 `json:"percentage"`
				PointsAgainst int64   `json:"pointsAgainst"`
				PointsFor     int64   `json:"pointsFor"`
				StreakLength  int64   `json:"streakLength"`
				StreakType    string  `json:"streakType"`
				Ties          int64   `json:"ties"`
				Wins          int64   `json:"wins"`
			} `json:"division"`
			Home struct {
				GamesBack     int64  `json:"gamesBack"`
				Losses        int64  `json:"losses"`
				Percentage    int64  `json:"percentage"`
				PointsAgainst int64  `json:"pointsAgainst"`
				PointsFor     int64  `json:"pointsFor"`
				StreakLength  int64  `json:"streakLength"`
				StreakType    string `json:"streakType"`
				Ties          int64  `json:"ties"`
				Wins          int64  `json:"wins"`
			} `json:"home"`
			Overall struct {
				GamesBack     int64   `json:"gamesBack"`
				Losses        int64   `json:"losses"`
				Percentage    float64 `json:"percentage"`
				PointsAgainst float64 `json:"pointsAgainst"`
				PointsFor     float64 `json:"pointsFor"`
				StreakLength  int64   `json:"streakLength"`
				StreakType    string  `json:"streakType"`
				Ties          int64   `json:"ties"`
				Wins          int64   `json:"wins"`
			} `json:"overall"`
		} `json:"record"`
		Roster struct {
			AppliedStatTotal float64 `json:"appliedStatTotal"`
			Entries          []struct {
				AcquisitionDate       int64       `json:"acquisitionDate"`
				AcquisitionType       string      `json:"acquisitionType"`
				InjuryStatus          string      `json:"injuryStatus"`
				LineupSlotID          int64       `json:"lineupSlotId"`
				PendingTransactionIds interface{} `json:"pendingTransactionIds"`
				PlayerID              int64       `json:"playerId"`
				PlayerPoolEntry       struct {
					AppliedStatTotal  int64 `json:"appliedStatTotal"`
					ID                int64 `json:"id"`
					KeeperValue       int64 `json:"keeperValue"`
					KeeperValueFuture int64 `json:"keeperValueFuture"`
					LineupLocked      bool  `json:"lineupLocked"`
					OnTeamID          int64 `json:"onTeamId"`
					Player            struct {
						Active               bool  `json:"active"`
						DefaultPositionID    int64 `json:"defaultPositionId"`
						DraftRanksByRankType struct {
							Standard struct {
								AuctionValue int64  `json:"auctionValue"`
								Rank         int64  `json:"rank"`
								RankSourceID int64  `json:"rankSourceId"`
								RankType     string `json:"rankType"`
								SlotID       int64  `json:"slotId"`
							} `json:"STANDARD"`
						} `json:"draftRanksByRankType"`
						Droppable     bool    `json:"droppable"`
						EligibleSlots []int64 `json:"eligibleSlots"`
						FirstName     string  `json:"firstName"`
						FullName      string  `json:"fullName"`
						ID            int64   `json:"id"`
						Injured       bool    `json:"injured"`
						InjuryStatus  string  `json:"injuryStatus"`
						Jersey        string  `json:"jersey"`
						LastName      string  `json:"lastName"`
						LastNewsDate  int64   `json:"lastNewsDate"`
						LastVideoDate int64   `json:"lastVideoDate"`
						Ownership     struct {
							AuctionValueAverage  float64 `json:"auctionValueAverage"`
							AverageDraftPosition int64   `json:"averageDraftPosition"`
							PercentChange        float64 `json:"percentChange"`
							PercentOwned         float64 `json:"percentOwned"`
							PercentStarted       float64 `json:"percentStarted"`
						} `json:"ownership"`
						ProTeamID int64 `json:"proTeamId"`
						Rankings  struct {
							Zero []struct {
								AuctionValue int64  `json:"auctionValue"`
								Rank         int64  `json:"rank"`
								RankSourceID int64  `json:"rankSourceId"`
								RankType     string `json:"rankType"`
								SlotID       int64  `json:"slotId"`
							} `json:"0"`
							One []struct {
								AuctionValue int64  `json:"auctionValue"`
								Rank         int64  `json:"rank"`
								RankSourceID int64  `json:"rankSourceId"`
								RankType     string `json:"rankType"`
								SlotID       int64  `json:"slotId"`
							} `json:"1"`
							One0 []struct {
								AuctionValue int64  `json:"auctionValue"`
								Rank         int64  `json:"rank"`
								RankSourceID int64  `json:"rankSourceId"`
								RankType     string `json:"rankType"`
								SlotID       int64  `json:"slotId"`
							} `json:"10"`
							One1 []struct {
								AuctionValue int64  `json:"auctionValue"`
								Rank         int64  `json:"rank"`
								RankSourceID int64  `json:"rankSourceId"`
								RankType     string `json:"rankType"`
								SlotID       int64  `json:"slotId"`
							} `json:"11"`
							One2 []struct {
								AuctionValue int64  `json:"auctionValue"`
								Rank         int64  `json:"rank"`
								RankSourceID int64  `json:"rankSourceId"`
								RankType     string `json:"rankType"`
								SlotID       int64  `json:"slotId"`
							} `json:"12"`
							One3 []struct {
								AuctionValue int64  `json:"auctionValue"`
								Rank         int64  `json:"rank"`
								RankSourceID int64  `json:"rankSourceId"`
								RankType     string `json:"rankType"`
								SlotID       int64  `json:"slotId"`
							} `json:"13"`
							One4 []struct {
								AuctionValue int64  `json:"auctionValue"`
								Rank         int64  `json:"rank"`
								RankSourceID int64  `json:"rankSourceId"`
								RankType     string `json:"rankType"`
								SlotID       int64  `json:"slotId"`
							} `json:"14"`
							One5 []struct {
								AuctionValue int64  `json:"auctionValue"`
								Rank         int64  `json:"rank"`
								RankSourceID int64  `json:"rankSourceId"`
								RankType     string `json:"rankType"`
								SlotID       int64  `json:"slotId"`
							} `json:"15"`
							One6 []struct {
								AuctionValue int64  `json:"auctionValue"`
								Rank         int64  `json:"rank"`
								RankSourceID int64  `json:"rankSourceId"`
								RankType     string `json:"rankType"`
								SlotID       int64  `json:"slotId"`
							} `json:"16"`
							One7 []struct {
								AuctionValue int64  `json:"auctionValue"`
								Rank         int64  `json:"rank"`
								RankSourceID int64  `json:"rankSourceId"`
								RankType     string `json:"rankType"`
								SlotID       int64  `json:"slotId"`
							} `json:"17"`
							Two []struct {
								AuctionValue int64  `json:"auctionValue"`
								Rank         int64  `json:"rank"`
								RankSourceID int64  `json:"rankSourceId"`
								RankType     string `json:"rankType"`
								SlotID       int64  `json:"slotId"`
							} `json:"2"`
							Three []struct {
								AuctionValue int64  `json:"auctionValue"`
								Rank         int64  `json:"rank"`
								RankSourceID int64  `json:"rankSourceId"`
								RankType     string `json:"rankType"`
								SlotID       int64  `json:"slotId"`
							} `json:"3"`
							Four []struct {
								AuctionValue int64  `json:"auctionValue"`
								Rank         int64  `json:"rank"`
								RankSourceID int64  `json:"rankSourceId"`
								RankType     string `json:"rankType"`
								SlotID       int64  `json:"slotId"`
							} `json:"4"`
							Five []struct {
								AuctionValue int64  `json:"auctionValue"`
								Rank         int64  `json:"rank"`
								RankSourceID int64  `json:"rankSourceId"`
								RankType     string `json:"rankType"`
								SlotID       int64  `json:"slotId"`
							} `json:"5"`
							Six []struct {
								AuctionValue int64  `json:"auctionValue"`
								Rank         int64  `json:"rank"`
								RankSourceID int64  `json:"rankSourceId"`
								RankType     string `json:"rankType"`
								SlotID       int64  `json:"slotId"`
							} `json:"6"`
							Seven []struct {
								AuctionValue int64  `json:"auctionValue"`
								Rank         int64  `json:"rank"`
								RankSourceID int64  `json:"rankSourceId"`
								RankType     string `json:"rankType"`
								SlotID       int64  `json:"slotId"`
							} `json:"7"`
							Eight []struct {
								AuctionValue int64  `json:"auctionValue"`
								Rank         int64  `json:"rank"`
								RankSourceID int64  `json:"rankSourceId"`
								RankType     string `json:"rankType"`
								SlotID       int64  `json:"slotId"`
							} `json:"8"`
							Nine []struct {
								AuctionValue int64  `json:"auctionValue"`
								Rank         int64  `json:"rank"`
								RankSourceID int64  `json:"rankSourceId"`
								RankType     string `json:"rankType"`
								SlotID       int64  `json:"slotId"`
							} `json:"9"`
						} `json:"rankings"`
						Stats []struct {
							AppliedAverage float64 `json:"appliedAverage"`
							AppliedStats   struct {
								One01  float64 `json:"101"`
								One02  float64 `json:"102"`
								One03  float64 `json:"103"`
								One04  float64 `json:"104"`
								One23  int64   `json:"123"`
								One24  int64   `json:"124"`
								One25  int64   `json:"125"`
								One28  int64   `json:"128"`
								One29  int64   `json:"129"`
								One30  int64   `json:"130"`
								One32  int64   `json:"132"`
								One33  int64   `json:"133"`
								One34  int64   `json:"134"`
								One35  int64   `json:"135"`
								One36  int64   `json:"136"`
								One9   float64 `json:"19"`
								Two0   float64 `json:"20"`
								Two06  int64   `json:"206"`
								Two4   float64 `json:"24"`
								Two5   float64 `json:"25"`
								Two6   float64 `json:"26"`
								Three  float64 `json:"3"`
								Four   int64   `json:"4"`
								Four2  float64 `json:"42"`
								Four3  float64 `json:"43"`
								Four4  float64 `json:"44"`
								Five3  float64 `json:"53"`
								Six3   float64 `json:"63"`
								Seven2 float64 `json:"72"`
								Seven4 float64 `json:"74"`
								Seven7 float64 `json:"77"`
								Eight0 float64 `json:"80"`
								Eight5 float64 `json:"85"`
								Eight6 float64 `json:"86"`
								Eight9 int64   `json:"89"`
								Nine0  int64   `json:"90"`
								Nine1  int64   `json:"91"`
								Nine2  int64   `json:"92"`
								Nine3  float64 `json:"93"`
								Nine5  float64 `json:"95"`
								Nine6  float64 `json:"96"`
								Nine7  float64 `json:"97"`
								Nine8  float64 `json:"98"`
								Nine9  float64 `json:"99"`
							} `json:"appliedStats"`
							AppliedTotal   float64 `json:"appliedTotal"`
							ExternalID     string  `json:"externalId"`
							ID             string  `json:"id"`
							LastUpdateInfo struct {
								ClientAddress interface{} `json:"clientAddress"`
								Platform      interface{} `json:"platform"`
								Source        interface{} `json:"source"`
							} `json:"lastUpdateInfo"`
							ProTeamID       int64 `json:"proTeamId"`
							ScoringPeriodID int64 `json:"scoringPeriodId"`
							SeasonID        int64 `json:"seasonId"`
							StatSourceID    int64 `json:"statSourceId"`
							StatSplitTypeID int64 `json:"statSplitTypeId"`
							Stats           struct {
								Zero   int64   `json:"0"`
								One    int64   `json:"1"`
								One0   int64   `json:"10"`
								One00  int64   `json:"100"`
								One01  float64 `json:"101"`
								One02  float64 `json:"102"`
								One03  float64 `json:"103"`
								One04  float64 `json:"104"`
								One05  float64 `json:"105"`
								One06  float64 `json:"106"`
								One07  int64   `json:"107"`
								One08  int64   `json:"108"`
								One09  int64   `json:"109"`
								One1   int64   `json:"11"`
								One10  int64   `json:"110"`
								One11  int64   `json:"111"`
								One12  float64 `json:"112"`
								One13  float64 `json:"113"`
								One14  int64   `json:"114"`
								One15  float64 `json:"115"`
								One16  int64   `json:"116"`
								One17  int64   `json:"117"`
								One18  int64   `json:"118"`
								One19  int64   `json:"119"`
								One2   int64   `json:"12"`
								One20  float64 `json:"120"`
								One21  int64   `json:"121"`
								One22  int64   `json:"122"`
								One23  int64   `json:"123"`
								One24  int64   `json:"124"`
								One25  int64   `json:"125"`
								One26  float64 `json:"126"`
								One27  float64 `json:"127"`
								One28  int64   `json:"128"`
								One29  int64   `json:"129"`
								One3   int64   `json:"13"`
								One30  int64   `json:"130"`
								One31  int64   `json:"131"`
								One32  int64   `json:"132"`
								One33  int64   `json:"133"`
								One34  int64   `json:"134"`
								One35  int64   `json:"135"`
								One36  int64   `json:"136"`
								One37  float64 `json:"137"`
								One38  int64   `json:"138"`
								One39  int64   `json:"139"`
								One4   int64   `json:"14"`
								One40  float64 `json:"140"`
								One41  float64 `json:"141"`
								One42  float64 `json:"142"`
								One43  int64   `json:"143"`
								One44  int64   `json:"144"`
								One45  float64 `json:"145"`
								One46  float64 `json:"146"`
								One47  int64   `json:"147"`
								One5   int64   `json:"15"`
								One54  int64   `json:"154"`
								One55  int64   `json:"155"`
								One56  int64   `json:"156"`
								One57  int64   `json:"157"`
								One6   int64   `json:"16"`
								One60  float64 `json:"160"`
								One7   float64 `json:"17"`
								One74  float64 `json:"174"`
								One8   float64 `json:"18"`
								One9   float64 `json:"19"`
								One98  float64 `json:"198"`
								One99  float64 `json:"199"`
								Two    float64 `json:"2"`
								Two0   float64 `json:"20"`
								Two00  float64 `json:"200"`
								Two05  int64   `json:"205"`
								Two06  int64   `json:"206"`
								Two1   int64   `json:"21"`
								Two10  int64   `json:"210"`
								Two2   float64 `json:"22"`
								Two3   float64 `json:"23"`
								Two4   float64 `json:"24"`
								Two5   float64 `json:"25"`
								Two6   float64 `json:"26"`
								Two7   int64   `json:"27"`
								Two8   int64   `json:"28"`
								Two9   int64   `json:"29"`
								Three  int64   `json:"3"`
								Three0 int64   `json:"30"`
								Three1 int64   `json:"31"`
								Three2 int64   `json:"32"`
								Three3 int64   `json:"33"`
								Three4 int64   `json:"34"`
								Three5 float64 `json:"35"`
								Three6 float64 `json:"36"`
								Three7 float64 `json:"37"`
								Three8 float64 `json:"38"`
								Three9 float64 `json:"39"`
								Four   int64   `json:"4"`
								Four0  float64 `json:"40"`
								Four1  float64 `json:"41"`
								Four2  float64 `json:"42"`
								Four3  float64 `json:"43"`
								Four4  float64 `json:"44"`
								Four5  float64 `json:"45"`
								Four6  float64 `json:"46"`
								Four7  int64   `json:"47"`
								Four8  int64   `json:"48"`
								Four9  int64   `json:"49"`
								Five   int64   `json:"5"`
								Five0  int64   `json:"50"`
								Five1  int64   `json:"51"`
								Five2  int64   `json:"52"`
								Five3  float64 `json:"53"`
								Five4  int64   `json:"54"`
								Five5  int64   `json:"55"`
								Five6  float64 `json:"56"`
								Five7  float64 `json:"57"`
								Five8  float64 `json:"58"`
								Five9  int64   `json:"59"`
								Six    int64   `json:"6"`
								Six0   float64 `json:"60"`
								Six1   float64 `json:"61"`
								Six2   float64 `json:"62"`
								Six3   float64 `json:"63"`
								Six4   float64 `json:"64"`
								Six5   float64 `json:"65"`
								Six6   float64 `json:"66"`
								Six7   float64 `json:"67"`
								Six8   float64 `json:"68"`
								Six9   float64 `json:"69"`
								Seven  int64   `json:"7"`
								Seven0 float64 `json:"70"`
								Seven1 float64 `json:"71"`
								Seven2 float64 `json:"72"`
								Seven3 float64 `json:"73"`
								Seven4 float64 `json:"74"`
								Seven5 float64 `json:"75"`
								Seven6 float64 `json:"76"`
								Seven7 float64 `json:"77"`
								Seven8 float64 `json:"78"`
								Seven9 float64 `json:"79"`
								Eight  int64   `json:"8"`
								Eight0 float64 `json:"80"`
								Eight1 float64 `json:"81"`
								Eight2 float64 `json:"82"`
								Eight3 float64 `json:"83"`
								Eight4 float64 `json:"84"`
								Eight5 float64 `json:"85"`
								Eight6 float64 `json:"86"`
								Eight7 float64 `json:"87"`
								Eight8 float64 `json:"88"`
								Eight9 int64   `json:"89"`
								Nine   int64   `json:"9"`
								Nine0  int64   `json:"90"`
								Nine1  int64   `json:"91"`
								Nine2  int64   `json:"92"`
								Nine3  float64 `json:"93"`
								Nine4  float64 `json:"94"`
								Nine5  float64 `json:"95"`
								Nine6  float64 `json:"96"`
								Nine7  float64 `json:"97"`
								Nine8  float64 `json:"98"`
								Nine9  float64 `json:"99"`
							} `json:"stats"`
						} `json:"stats"`
						UniverseID int64 `json:"universeId"`
					} `json:"player"`
					Ratings struct {
						Zero struct {
							PositionalRanking int64   `json:"positionalRanking"`
							TotalRanking      int64   `json:"totalRanking"`
							TotalRating       float64 `json:"totalRating"`
						} `json:"0"`
					} `json:"ratings"`
					RosterLocked bool   `json:"rosterLocked"`
					Status       string `json:"status"`
					TradeLocked  bool   `json:"tradeLocked"`
				} `json:"playerPoolEntry"`
				Status string `json:"status"`
			} `json:"entries"`
		} `json:"roster"`
		TradeBlock         struct{} `json:"tradeBlock"`
		TransactionCounter struct {
			AcquisitionBudgetSpent   int64 `json:"acquisitionBudgetSpent"`
			Acquisitions             int64 `json:"acquisitions"`
			Drops                    int64 `json:"drops"`
			MatchupAcquisitionTotals struct {
				One   int64 `json:"1"`
				One0  int64 `json:"10"`
				One1  int64 `json:"11"`
				One2  int64 `json:"12"`
				One3  int64 `json:"13"`
				One4  int64 `json:"14"`
				One5  int64 `json:"15"`
				Two   int64 `json:"2"`
				Three int64 `json:"3"`
				Four  int64 `json:"4"`
				Five  int64 `json:"5"`
				Six   int64 `json:"6"`
				Seven int64 `json:"7"`
				Eight int64 `json:"8"`
				Nine  int64 `json:"9"`
			} `json:"matchupAcquisitionTotals"`
			Misc         int64 `json:"misc"`
			MoveToActive int64 `json:"moveToActive"`
			MoveToIR     int64 `json:"moveToIR"`
			Paid         int64 `json:"paid"`
			TeamCharges  int64 `json:"teamCharges"`
			Trades       int64 `json:"trades"`
		} `json:"transactionCounter"`
		WaiverRank int64   `json:"waiverRank"`
		WatchList  []int64 `json:"watchList"`
	} `json:"teams"`
}
