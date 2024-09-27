#define _CRT_SECURE_NO_WARNINGS

#include<SDL.h>
#include<SDL_ttf.h>
#include<SDL_mixer.h>
#include<SDL_image.h>
#include<stdbool.h>
#include<stdio.h>
#include<stdlib.h>
#include<time.h>

#define MAXLEN 500
//窗口宽、高度
const int Screen_width = 640;
const int Screen_height = 480;

/*
	1.包含目录：头文件
	2.库目录：  lib
	3.链接器->输入->附加依赖项：所有的lib名字加进来
	4.DLL path路径配置 <- 环境变量
	5.导出为模板
*/

//贪吃蛇
typedef struct Snake {
	SDL_Point pos[MAXLEN];//蛇的身体坐标数组
	int size;//宽度和高度
	int len;//蛇当前长度
	int dir;//蛇的移动方向
}Snake;
Snake snake;

//果子
typedef struct Fruit {
	SDL_Point pos_f;
}Fruit;

Fruit fruit;

void fruit_a(SDL_Point  point , SDL_Renderer* render) {
	SDL_SetRenderDrawColor(render, 0, 0, 255, 255);
	
	SDL_Rect rect_f = { point.x,point.y,10,10 };
	SDL_RenderDrawRect(render, &rect_f);
}


void snake_init(Snake* snake,int len) {
	snake->len = len;
	snake->size = 10;
	snake->dir = SDLK_RIGHT;
	for (int i = 0; i < snake->len; i++) {
		snake->pos[i].x = (snake->len - i - 1) * snake->size;//也可更改循环条件
		snake->pos[i].y = 0;
	}

}
void snake_draw(Snake* snake,SDL_Renderer* render) {
	
	SDL_SetRenderDrawColor(render, 255, 0, 0, 255);
	for (int i = 0; i < snake->len; i++) {
		//绘制矩形
		SDL_Rect rect = { snake->pos[i].x,snake->pos[i].y,snake->size,snake->size};
		SDL_RenderDrawRect(render, &rect);
	}
}
//让坐标发生变化即可
void snake_move(Snake* snake) {
	//只要让蛇头变即可，身体是跟着蛇头走的
	for (int i = snake->len - 1; i > 0; i--) //有bug
	{
		snake->pos[i] = snake->pos[i - 1];
	}
	//让蛇改变方向
	switch (snake->dir) {
	case SDLK_UP: //↑，下方以此类推
		snake->pos[0].y -= snake->size;
		break;
	case SDLK_DOWN:
		snake->pos[0].y += snake->size;
		break;
	case SDLK_LEFT:
		snake->pos[0].x -= snake->size;
		break;
	case SDLK_RIGHT:
		snake->pos[0].x += snake->size;
		break;
	}
}

SDL_Point random(SDL_Point point){
	point.x = 0;
	point.y= 0;
	int m, n;
	time_t t;
	srand((unsigned)time(&t));

	m = rand() % 62 + 1;
	n = rand() % 46 + 1;
	point.x = m * 10;
	point.y = n * 10;
    return point;
}

void draw(SDL_Renderer* render) {
	//设置窗口背景颜色
	SDL_SetRenderDrawColor(render, 255, 255, 255, 255);//RGBA模式    α 0~255 0:完全透明
	//用当前颜色填充整个屏幕
	SDL_RenderClear(render);

	SDL_SetRenderDrawColor(render, 200, 200, 200, 255);

	//画网格
	for (size_t i = 0; i < Screen_height; i++) {
		SDL_RenderDrawLine(render, 0, i * 10, Screen_width, i * 10);
	}
	for (size_t i = 0; i < Screen_width; i++) {
		SDL_RenderDrawLine(render, i * 10, 0, i * 10, Screen_height);
	}
	//绘制蛇

	snake_draw(&snake, render);

	//把渲染器数据全部绘制到窗口上
	SDL_RenderPresent(render);
}

void drawFruit(SDL_Renderer* render) {
	fruit_a(fruit.pos_f, render);

	//把渲染器数据全部绘制到窗口上
	SDL_RenderPresent(render);
}

void init_sdl() 
{
	snake_init(&snake,4);
}
//void studyDataType() {
//	SDL_Point point;
//	SDL_Rect rect;
//	SDL_Color color;//RGBA 含有alpha通道 —— 透明度
//
//	SDL_Window* window = NULL;
//	SDL_Renderer* render = NULL;
//	SDL_Surface* Surface = NULL;
//	SDL_Texture* texture = NULL;
//	
//
//	SDL_Event msg;
//	//存储鼠标的x,y坐标
//	msg.motion.x;
//	msg.motion.y; 
//
//	//存储键盘按下的键
//	msg.key.keysym.sym;
//	
//	//存储鼠标按下了哪个键
//	msg.button.button;
//
//}
//void createWindow(int width,int height) {
//	SDL_Window* w = SDL_CreateWindow("cool", 100, 100, width, height, 4);
//}

Uint32* snakeMoveCall(Uint32 interval, void* param) {
	snake_move(&snake);
	return interval;
}

//处理键盘按下事件
bool keyPressEventt(SDL_KeyboardEvent* ev) {
	if(ev->keysym.sym == SDLK_SPACE)
	{ 	return true;
	}
	return false;
}


//处理键盘按下事件
void keyPressEvent(SDL_KeyboardEvent * ev) {
	switch (ev->keysym.sym)
	{
	case SDLK_UP: //↑，下方以此类推
	/*	SDL_Log("up\n");*/
		if (snake.dir == SDLK_DOWN)
			break;
		snake.dir = SDLK_UP;
		break;
	case SDLK_DOWN:
	/*	SDL_Log("Down\n");*/
		if (snake.dir == SDLK_UP) {
			snake.dir == SDLK_UP;
			break;
		}
		snake.dir = SDLK_DOWN;
		break;
	case SDLK_LEFT:
	/*	SDL_Log("Left\n");*/
		if (snake.dir == SDLK_RIGHT) {
			snake.dir == SDLK_RIGHT;
				break;
		}
		snake.dir = SDLK_LEFT;
		break;
	case SDLK_RIGHT:
	/*	SDL_Log("Right\n");*/
		if (snake.dir == SDLK_LEFT) {
			snake.dir == SDLK_LEFT;
				break;
		}
		snake.dir = SDLK_RIGHT;
		break;
	}
}
int showone() {
	// 显示一个带有两个按钮的信息框
	SDL_MessageBoxButtonData buttons[] = {
		{ SDL_MESSAGEBOX_BUTTON_RETURNKEY_DEFAULT, 1, "Quit" },
		{ SDL_MESSAGEBOX_BUTTON_ESCAPEKEY_DEFAULT, 2, "Retry" }
	};

	SDL_MessageBoxData messageboxdata = {
		SDL_MESSAGEBOX_INFORMATION,  // 类型
		NULL,                        // 窗口
		"\xe6\x8f\x90\xe7\xa4\xba",                      // 标题
		"\xe8\xaf\xb7\xe9\x80\x89\xe6\x8b\xa9\xe6\x93\x8d\xe4\xbd\x9c",                 // 内容
		SDL_arraysize(buttons),       // 按钮数量
		buttons,                     // 按钮数据
		NULL                          // 默认按钮
	};

	int buttonid;
	SDL_ShowMessageBox(&messageboxdata, &buttonid);

	// 根据用户选择的按钮执行相应操作
	return buttonid;
}

int main(int argc, char* argv[])
{
	//SDL_Init(SDL_EVERYTHING)
	if (SDL_Init(SDL_INIT_VIDEO|SDL_INIT_TIMER|SDL_INIT_AUDIO)) {
		SDL_Log("%s", SDL_GetError());
		return -1;
	}
	TTF_Init();
	IMG_Init(IMG_INIT_JPG);

	Mix_OpenAudio(44100, MIX_DEFAULT_FORMAT, 2, 2048);
	//创建窗口和渲染器(一块内存，暂存绘图信息，双缓冲)
	SDL_Window* window = NULL;//窗口指针
	SDL_Renderer* render = NULL;//渲染器指针
	SDL_CreateWindowAndRenderer(Screen_width, Screen_height, 4, &window,&render);
	SDL_SetWindowTitle(window,"\xe8\xb4\xaa\xe5\x90\x83\xe8\x9b\x87(\xe7\xbb\x8f\xe5\x85\xb8\xe6\xa8\xa1\xe5\xbc\x8f)");//使用python的 encode函数转换为utf8格式
	SDL_Surface* icon = SDL_LoadBMP("xiaohui.bmp");
	SDL_SetWindowIcon(window, icon);
	
	SDL_Surface* imageSurface = IMG_Load("background.bmp");
	SDL_Texture* imageTexture = SDL_CreateTextureFromSurface(render, imageSurface);
	SDL_FreeSurface(imageSurface);


	SDL_Event event;
	bool isdonee = true;
	while (isdonee) {
		//	SDL_RenderCopy(render, imageTexture, NULL, NULL);
		SDL_RenderPresent(render);
		//
		while (SDL_PollEvent(&event)) {
			if (event.type == SDL_MOUSEBUTTONDOWN || (event.type == SDL_KEYDOWN && keyPressEventt(&event.key))) {
				// 清空渲染器
				SDL_RenderClear(render);

				// 绘制第二个界面的内容
				isdonee = false;
				draw(render);
				drawFruit(render);

				SDL_RenderPresent(render);
			}
		}
	}

	SDL_AddTimer(200, snakeMoveCall, NULL); // 定时器
a:
	init_sdl();

	fruit.pos_f = random(fruit.pos_f);
	
	TTF_Font* font = TTF_OpenFont("ERASMD.ttf", 24);
	SDL_Color textColor = { 0,0,0 };
	
	// 渲染分数的纹理
	SDL_Texture* scoreTexture = NULL;


	Mix_Music* bgm = Mix_LoadMUS("background_music.mp3");  // 加载背景音乐文件
	if (bgm == NULL) {
		return 1;
	}
	
	Mix_Chunk* soundEffect = Mix_LoadWAV("sound_effect.wav");
	if (soundEffect == NULL) {
		// 处理加载失败的情况
		return 1;
	}
	// 播放背景音乐
	Mix_PlayMusic(bgm, -1);  // -1表示循环播放
	//创建一个定时器 原位置
	//事件处理



	bool isdone = false;
	while (!isdone) {
		//帧率   fps 最低24帧  一般60帧   --每秒绘图的次数
		//获取当前时间
		Uint32 begTime = SDL_GetTicks(); //获取现在毫秒数

		draw(render);
		drawFruit(render);

		//定义事件变量
		SDL_Event ev;

		//获取事件,在一次循环中 把所有事件都处理完成
		while (SDL_PollEvent(&ev)) {
			switch (ev.type) {
			case SDL_QUIT:  //关闭窗口事件
				isdone = true;
				break;
			case SDL_KEYDOWN://按键按下
				keyPressEvent(&ev.key);
				break;
			}
		}
		char scoreText[50];

		sprintf(scoreText, "Score: %d", snake.len - 4);
		if (scoreTexture) {
			SDL_DestroyTexture(scoreTexture);
		}
		SDL_Surface* surface_1 = TTF_RenderText_Solid(font, scoreText, textColor);

		scoreTexture = SDL_CreateTextureFromSurface(render, surface_1);
		SDL_Rect textRect = { 10, 10, surface_1->w, surface_1->h };

		SDL_RenderCopy(render, scoreTexture, NULL, &textRect);
		// 更新窗口
		SDL_RenderPresent(render);
		// 释放表面和纹理
		SDL_FreeSurface(surface_1);
		if (scoreTexture) {
			SDL_DestroyTexture(scoreTexture);
		}
		if (fruit.pos_f.x == snake.pos[0].x && fruit.pos_f.y == snake.pos[0].y)
		{
			SDL_Log("%d,%d,%d", snake.pos[0].x, snake.pos[0].y, snake.len + 1);
			fruit.pos_f = random(fruit.pos_f);
			snake.len++;
		}

		if ((snake.pos[0].x == 630 && snake.dir == SDLK_RIGHT) || (snake.pos[0].y == 470 && snake.dir == SDLK_DOWN) || (snake.pos[0].x == 0 && snake.dir == SDLK_LEFT) || (snake.pos[0].y == 0 && snake.dir == SDLK_UP)) {
			Mix_FreeMusic(bgm);  // 释放背景音乐资源
			Mix_PlayChannel(-1, soundEffect, 0);  // -1表示使用第一个可用的通道，0表示播放一次
			switch (showone()) {
			case 1:  // Quit
				SDL_Quit();
				break;
			case 2:  // Retry
				goto a;
				break;
			default:
				// 其他操作
				break;
			}
		}
		if (snake.len > 4) {
			for (int i = 1; i < snake.len; i++) {
				if (snake.pos[0].x == snake.pos[i].x && snake.pos[0].y == snake.pos[i].y) {
					Mix_FreeMusic(bgm);  // 释放背景音乐资源
					Mix_PlayChannel(-1, soundEffect, 0);  // -1表示使用第一个可用的通道，0表示播放一次
					/*if (SDL_ShowSimpleMessageBox(SDL_MESSAGEBOX_INFORMATION, "bad news !", "you lose", window)) {
						SDL_Log("%s", SDL_GetError());
						return -1;
					}*/

					switch (showone()) {
					case 1:  // Quit
						SDL_Quit();
						break;
					case 2:  // Retry
						goto a;
						break;
					default:
						// 其他操作
						break;
					}
				}
			}
		}
		//控制帧率
		Uint32 eaplsed = SDL_GetTicks() - begTime; //经过时间
		int delay = 1000 / 60.0 - eaplsed;
		if (delay > 0) SDL_Delay(delay);//动态延迟
	}

	Mix_FreeMusic(bgm);  // 释放背景音乐资源
	Mix_FreeChunk(soundEffect);
	Mix_CloseAudio();  // 关闭音频设备
	//动态内存的释放
	SDL_FreeSurface(icon);
	TTF_CloseFont(font);
	SDL_DestroyWindow(window);
	SDL_DestroyRenderer(render);
	TTF_Quit();
	SDL_Quit();
	return 0;
}