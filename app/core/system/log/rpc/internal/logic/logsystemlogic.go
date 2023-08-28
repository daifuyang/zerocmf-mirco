package logic

import (
	"app/core/log/model"
	"context"
	"github.com/zeromicro/go-zero/core/logc"
	"strings"

	"app/core/log/rpc/internal/svc"
	"app/core/log/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogSystemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLogSystemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogSystemLogic {
	return &LogSystemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 系统日志记录
func (l *LogSystemLogic) LogSystem(in *pb.SystemLogReq) (*pb.SystemLogResp, error) {
	systemLogModel := l.svcCtx.SystemLogModel

	appName := in.GetAppName()
	if strings.TrimSpace(appName) == "" {
		logc.Error(context.Background(), "服务名称为空！")
	}

	loglevel := in.GetLogLevel()
	if strings.TrimSpace(loglevel) == "" {
		logc.Error(context.Background(), "日志等级为空！")
	}

	method := in.GetMethod()
	if strings.TrimSpace(method) == "" {
		logc.Error(context.Background(), "日志方法为空！")
	}

	message := in.GetLogLevel()
	if strings.TrimSpace(message) == "" {
		logc.Error(context.Background(), "日志消息为空！")
	}

	systemLog := model.SystemLog{
		AppName:  appName,
		LogLevel: loglevel,
		Method:   method,
		Message:  message,
	}

	_, err := systemLogModel.Insert(l.ctx, &systemLog)
	if err != nil {
		return nil, err
	}

	return &pb.SystemLogResp{
		Success: true,
		Message: "新增成功！",
	}, nil
}
