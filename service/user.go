package service

import (
	"github.com/PedroGao/shoot/model"
	"github.com/PedroGao/shoot/utils"
	"sync"
)

func ListUser() ([]*model.User, error) {
	// 1.新建存放用户信息的数组
	infos := make([]*model.User, 0)
	// 2.从数据库查询用户信息
	users := []*model.User{
		{
			Id:       1,
			Nickname: "pedro",
			Password: "123456",
		},
		{
			Id:       2,
			Nickname: "pedro1",
			Password: "123456",
		},
		{
			Id:       3,
			Nickname: "pedro2",
			Password: "123456",
		},
	}
	// 3. 新建存放用户id的数组
	var ids []int
	//ids := make([]uint64, 0)
	for _, user := range users {
		ids = append(ids, user.Id)
	}

	// 4. 新建goroutine等待队列
	wg := sync.WaitGroup{}
	// 5. 新建用户列表映射
	userList := model.UserList{
		Lock:  new(sync.Mutex),
		IdMap: make(map[int]*model.User, len(users)),
	}

	// 6. 新建错误、完成通道
	errChan := make(chan error, 1)
	finished := make(chan bool, 1)

	// Improve query efficiency in parallel
	// 7. 并发处理数据
	for _, u := range users {
		wg.Add(1)
		go func(u *model.User) {
			defer wg.Done()

			shortId, err := utils.GenShortId()
			// 如果获取id错误，就写入错误通道
			if err != nil {
				errChan <- err
				return
			}

			//map是并发不安全的，因此需要加锁
			userList.Lock.Lock()
			defer userList.Lock.Unlock()
			// 以 key 为 id，value 为 userInfo填充map
			userList.IdMap[u.Id] = &model.User{
				Id:       u.Id,
				Nickname: u.Nickname + shortId,
				Password: u.Password,
			}
		}(u)
	}

	// 开启一个goroutine等待wg
	go func() {
		wg.Wait()
		//fmt.Println("close finished here!!!")
		//close(finished)
		finished <- true
	}()

	// 阻塞
	select {
	case <-finished:
		//fmt.Println("received finished here!!!")
	case err := <-errChan: // 当通道收到错误，就会立即退出当前函数，并返回错误
		return nil, err
	}

	for _, id := range ids {
		//fmt.Println("received an id , here")
		infos = append(infos, userList.IdMap[id])
	}

	return infos, nil
}
