package plugins

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/lib"
	"github.com/Dmitriy-Vas/wave/lib/packets/incoming"
	"github.com/Dmitriy-Vas/wave/lib/packets/outgoing"
	"log"
)

// AllAchievements claims all steam achievements after receiving ClientRevisionPacket
func AllAchievements(proxy *wave.Proxy) {
	proxy.HookPacket(int64(lib.OutClientRevision), true, func(conn *wave.Conn, packet wave.Packet) {
		if p := packet.(*outgoing.ClientRevisionPacket); !p.IsSteam {
			log.Printf("Please run game from the Steam to claim achievements")
			return
		}

		done := make(chan error)
		defer close(done)
		go func(done chan error, conn *wave.Conn) {
			for {
				err, closed := <-done
				if closed {
					log.Printf("All achievements claimed")
					break
				}
				if err != nil {
					conn.Close(err)
					break
				}
			}
		}(done, conn)

		p := &incoming.SteamAchievementPacket{
			ID:              int64(lib.IncSteamAchievement),
			Send:            true,
			Achievement:     lib.ACHIEVEMENT_AW_QUESTCOMPLETE,
			AchievementStat: lib.ACHIEVEMENT_STAT_QUESTCOUNT,
		}
		for i := 0; i <= 100; i++ {
			done <- conn.SendPacket(p, false)
		}

		p.Achievement = lib.ACHIEVEMENT_AW_GUMBYKILL100K
		p.AchievementStat = lib.ACHIEVEMENT_STAT_GUMBYCOUNT
		for i := 0; i <= 100000; i++ {
			done <- conn.SendPacket(p, false)
		}

		p.Achievement = lib.ACHIEVEMENT_AW_MONSTERKILL
		p.AchievementStat = lib.ACHIEVEMENT_STAT_KILLCOUNT
		for i := 0; i <= 100000; i++ {
			done <- conn.SendPacket(p, false)
		}

		p.Achievement = lib.ACHIEVEMENT_AW_PLAYERKILLER
		done <- conn.SendPacket(p, false)

		p.Achievement = lib.ACHIEVEMENT_AW_LICHCAPE
		done <- conn.SendPacket(p, false)

		p.Achievement = lib.ACHIEVEMENT_AW_HALLOWEENTOWER
		done <- conn.SendPacket(p, false)

		p.Achievement = lib.ACHIEVEMENT_AW_DEATH
		done <- conn.SendPacket(p, false)

		p.Achievement = lib.ACHIEVEMENT_AW_KILL_LUCAS
		done <- conn.SendPacket(p, false)

		p.Achievement = lib.ACHIEVEMENT_AW_ANNIVERSARY
		done <- conn.SendPacket(p, false)

		p.Achievement = lib.ACHIEVEMENT_AW_CREATECHAR
		done <- conn.SendPacket(p, false)

		p.Achievement = lib.ACHIEVEMENT_AW_BUYPREMIUMITEM
		done <- conn.SendPacket(p, false)

		p.Achievement = lib.ACHIEVEMENT_AW_ENHANCEITEM
		done <- conn.SendPacket(p, false)

		p.Achievement = lib.ACHIEVEMENT_AW_FIRSTMOUNT
		done <- conn.SendPacket(p, false)

		p.Achievement = lib.ACHIEVEMENT_AW_REACHESIA
		done <- conn.SendPacket(p, false)

		p.Achievement = lib.ACHIEVEMENT_AW_REACHGOLDUM
		done <- conn.SendPacket(p, false)

		p.Achievement = lib.ACHIEVEMENT_AW_REACHTERION
		done <- conn.SendPacket(p, false)

		p.Achievement = lib.ACHIEVEMENT_AW_REACHROKKUNTAVERN
		done <- conn.SendPacket(p, false)

		p.Achievement = lib.ACHIEVEMENT_AW_REACHCEMETERY
		done <- conn.SendPacket(p, false)
	})
}
