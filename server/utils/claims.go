package utils

import (
	"context"
	"net"
	"time"


	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/google/uuid"
)

func ClearToken(ctx context.Context, c *app.RequestContext) {
	// 增加cookie x-token 向来源的web添加
	host, _, err := net.SplitHostPort(string(c.Request.Host()))
	if err != nil {
		host = string(c.Request.Host())
	}

	if net.ParseIP(host) != nil {
		c.SetCookie("x-token", "", -1, "/", "", protocol.CookieSameSiteDefaultMode, false, false)
	} else {
		c.SetCookie("x-token", "", -1, "/", host, protocol.CookieSameSiteDefaultMode, false, false)
	}
}

func SetToken(c *app.RequestContext, token string, maxAge int) {
	// 增加cookie x-token 向来源的web添加
	host, _, err := net.SplitHostPort(string(c.Request.Host()))
	if err != nil {
		host = string(c.Request.Host())
	}

	if net.ParseIP(host) != nil {
		c.SetCookie("x-token", token, maxAge, "/", "", protocol.CookieSameSiteDefaultMode, false, false)
	} else {
		c.SetCookie("x-token", token, maxAge, "/", host, protocol.CookieSameSiteDefaultMode, false, false)
	}
}

func GetToken(ctx context.Context, c *app.RequestContext) string {
	token := string(c.Request.Header.Get("x-token"))
	if token == "" {
		j := NewJWT()
		token = string(c.Cookie("x-token"))
		claims, err := j.ParseToken(token)
		if err != nil {
			global.GVA_LOG.Error("重新写入cookie token失败,未能成功解析token,请检查请求头是否存在x-token且claims是否为规定结构")
			return token
		}
		SetToken(c, token, int(claims.ExpiresAt.Unix()-time.Now().Unix()))
	}
	return token
}

func GetClaims(ctx context.Context, c *app.RequestContext) (*systemReq.CustomClaims, error) {
	token := GetToken(ctx, c)
	j := NewJWT()
	claims, err := j.ParseToken(token)
	if err != nil {
		global.GVA_LOG.Error("从Gin的Context中获取从jwt解析信息失败, 请检查请求头是否存在x-token且claims是否为规定结构")
	}
	return claims, err
}

// GetUserID 从Gin的Context中获取从jwt解析出来的用户ID
func GetUserID(ctx context.Context, c *app.RequestContext) uint {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(ctx, c); err != nil {
			return 0
		} else {
			return cl.BaseClaims.ID
		}
	} else {
		waitUse := claims.(*systemReq.CustomClaims)
		return waitUse.BaseClaims.ID
	}
}

// GetUserUuid 从Gin的Context中获取从jwt解析出来的用户UUID
func GetUserUuid(ctx context.Context, c *app.RequestContext) uuid.UUID {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(ctx, c); err != nil {
			return uuid.UUID{}
		} else {
			return cl.UUID
		}
	} else {
		waitUse := claims.(*systemReq.CustomClaims)
		return waitUse.UUID
	}
}

// GetUserAuthorityId 从Gin的Context中获取从jwt解析出来的用户角色id
func GetUserAuthorityId(ctx context.Context, c *app.RequestContext) uint {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(ctx, c); err != nil {
			return 0
		} else {
			return cl.AuthorityId
		}
	} else {
		waitUse := claims.(*systemReq.CustomClaims)
		return waitUse.AuthorityId
	}
}

// GetUserInfo 从Gin的Context中获取从jwt解析出来的用户角色id
func GetUserInfo(ctx context.Context, c *app.RequestContext) *systemReq.CustomClaims {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(ctx, c); err != nil {
			return nil
		} else {
			return cl
		}
	} else {
		waitUse := claims.(*systemReq.CustomClaims)
		return waitUse
	}
}

// GetUserName 从Gin的Context中获取从jwt解析出来的用户名
func GetUserName(ctx context.Context, c *app.RequestContext) string {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(ctx, c); err != nil {
			return ""
		} else {
			return cl.Username
		}
	} else {
		waitUse := claims.(*systemReq.CustomClaims)
		return waitUse.Username
	}
}

func LoginToken(user system.Login) (token string, claims systemReq.CustomClaims, err error) {
	j := NewJWT()
	claims = j.CreateClaims(systemReq.BaseClaims{
		UUID:        user.GetUUID(),
		ID:          user.GetUserId(),
		NickName:    user.GetNickname(),
		Username:    user.GetUsername(),
		AuthorityId: user.GetAuthorityId(),
	})
	token, err = j.CreateToken(claims)
	return
}
