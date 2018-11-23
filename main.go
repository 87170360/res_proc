package main
import (
	"fmt"
	"os"
	"io"
	"strings"
	"path/filepath"
)

const (
    inputDir1 = "input1"
    inputDir2 = "input2"
    inputDir3 = "input3"
    outputDir = "output"
)

var (
	//不使用
	droplist = map[string]bool {
		"input2\\UI-美術資源繁體\\14.斗酒\\module33_doujiu_lose.png":true,
		"input2\\UI-美術資源繁體\\14.斗酒\\module33_doujiu_win.png":true,
		"input2\\UI-美術資源繁體\\14.斗酒\\斗酒\\module33_doujiu_guizejieshao.png":true,
		"input2\\UI-美術資源繁體\\23.押镖\\module21_yabiao_14.png":true,
		"input2\\UI-美術資源繁體\\13.抽奖\\抽奖\\module33_choujiang_jifenpaihangbang_zi.png":true,
		"input2\\UI-美術資源繁體\\15.斗老千\\斗老千\\module33_doulaoqian_jifengjiangli_biaoti.png":true,
		"input2\\UI-美術資源繁體\\15.斗老千\\斗老千\\图层-443111.png":true,
		"input2\\UI-美術資源繁體\\2.创建角色\\module02_role_huantouxiang_c.png": true,
	}

	//在处理ui_xxx_a.png -> a.png中，出现重复但不打印错误
	repeatForbitlist = map[string]bool {
		"input2\\UI-美術資源繁體\\5.角色（人物属性）\\人物境界\\1.png" : true,
		"input2\\UI-美術資源繁體\\5.角色（人物属性）\\人物境界\\2.png" : true,
		"input2\\UI-美術資源繁體\\5.角色（人物属性）\\人物境界\\3.png" : true,
		"input2\\UI-美術資源繁體\\5.角色（人物属性）\\人物境界\\4.png" : true,
		"input2\\UI-美術資源繁體\\5.角色（人物属性）\\人物境界\\5.png" : true,
		"input2\\UI-美術資源繁體\\5.角色（人物属性）\\人物境界\\6.png" : true,
		"input2\\UI-美術資源繁體\\5.角色（人物属性）\\人物境界\\7.png" : true,
		"input2\\UI-美術資源繁體\\5.角色（人物属性）\\人物境界\\8.png" : true,
	}

	//在处理ui_xxx_a.png -> a.png中，允许重复替换的
	repeatlist = map[string]bool {
		"input2\\UI-美術資源繁體\\29.江湖事\\module6_jianghushi_shuxing.png" : true,
		"input2\\UI-美術資源繁體\\27.活动\\1.签到\\签到\\module33_qiandao_yilingqu.png" : true,
		"input2\\UI-美術資源繁體\\0.共用（未完成）\\module14_wuxue_yizhuangbei.png" : true,
	}

	dirpro = []*dirprodata {
		&dirprodata{
			dir1 : "input3\\img\\icon\\jingjie",
			dir2 : "input2\\UI-美術資源繁體\\5.角色（人物属性）\\人物境界",
			//0.png -> jingjie_0.png 
			convert : func(str string) (fk, nn string) {
				fk = fmt.Sprintf("jingjie_%s", str)
				nn = fk
				return
			},
		},
		&dirprodata{
			dir1 : "input3\\img\\bg\\bg_jianghu",
			dir2 : "input2\\UI-美術資源繁體\\11.江湖\\输出文件",
			//0.png -> 0.png 
			convert : func(str string) (string, string) {
				return str, str
			},
		},
		&dirprodata{
			dir1 : "input3\\img\\bg\\bg_zlzc",
			dir2 : "input2\\UI-美術資源繁體\\30.逐鹿战场",
			//0.png -> 0.png 
			convert : func(str string) (string, string) {
				return str, str
			},
		},
		&dirprodata{
			dir1 : "input3\\img\\build",
			dir2 : "input2\\UI-美術資源繁體\\3.主界面\\功能建筑（繁体）",
			//0.png -> 0.png 
			convert : func(str string) (string, string) {
				return str, str
			},
		},
		&dirprodata{
			dir1 : "input3\\img\\bg\\bg_login",
			dir2 : "input2\\UI-美術資源繁體\\16.公告\\公告",
			//0.png -> 0.png 
			convert : func(str string) (string, string) {
				return str, str
			},
		},
	}

	filepro = []*fileprodata {
		&fileprodata{
			dir1 : "input1\\ui\\ui_major",
			dir2 : "input2\\UI-美術資源繁體\\11.江湖",
			m : map[string]string{
				"ui_major_unlock.png"			:"module6_jianghu_suo1.png",
				"ui_major_lock_2_hierarchy.png"	:"module6_jianghu_suo2.png",
				"ui_major_lock_2_level.png"		:"module6_jianghu_suo3.png",
				"ui_major_lock_5_level.png"		:"module6_jianghu_suo4.png",
				"ui_major_lock_1_hierarchy.png"	:"module6_jianghu_suo5.png",
				"ui_major_lock_4_level.png"		:"module6_jianghu_suo6.png",
				"ui_major_lock_3_level.png"		:"module6_jianghu_suo7.png",
				"ui_major_lock_1_level.png"		:"module6_jianghu_suo8.png",
				"ui_major_lock_3_hierarchy.png" :"module6_jianghu_suo9.png",
				"ui_major_lock_4_hierarchy.png" :"module6_jianghu_suo10.png",
				"ui_major_lock_5_hierarchy.png" :"module6_jianghu_suo11.png",
				"ui_major_lock_6_hierarchy.png" :"module6_jianghu_suo12.png",
			},
		},
		&fileprodata{
			dir1 : "input3\\img\\bg\\bg_jianghu",
			dir2 : "input2\\UI-美術資源繁體\\11.江湖\\输出文件",
			m : map[string]string{
				"module6_jianghu_zhuluzhanchang_a.png":"module6_jianghu_xiakexing_a.png",
				"module6_jianghu_zhuluzhanchang_b.png":"module6_jianghu_xiakexing_b.png",
			},
		},
		&fileprodata{
			dir1 : "input1\\ui\\ui_lilian",
			dir2 : "input2\\UI-美術資源繁體\\10.历练秘境",
			m : map[string]string{
				"ui_lilian_module06_lilian_mingcheng06.png":"module06_lilian _mc1.png",
				"ui_lilian_module06_lilian_mingcheng04.png":"module06_lilian _mc2.png",
				"ui_lilian_module06_lilian_mingcheng03.png":"module06_lilian _mc3.png",
				"ui_lilian_module06_lilian_mingcheng02.png":"module06_lilian _mc4.png",
				"ui_lilian_module06_lilian_mingcheng01.png":"module06_lilian _mc5.png",
				"ui_lilian_module06_lilian_mingcheng05.png":"module06_lilian _mc6.png",
			},
		},
		&fileprodata{
			dir1 : "input1\\ui\\ui_mijing",
			dir2 : "input2\\UI-美術資源繁體\\10.历练秘境",
			m : map[string]string{
				"ui_mijing_module06_mijing_tubiao01.png":"module06_mijing _tubiao01.png",
				"ui_mijing_module06_mijing_tubiao02.png":"module06_mijing _tubiao02.png",
				"ui_mijing_module06_mijing_tubiao03.png":"module06_mijing _tubiao03.png",
				"ui_mijing_module06_mijing_tubiao04.png":"module06_mijing _tubiao04.png",
				"ui_mijing_module06_mijing_tubiao05.png":"module06_mijing _tubiao05.png",
				"ui_mijing_module06_mijing_tubiao06.png":"module06_mijing _tubiao06.png",
			},
		},
		&fileprodata{
			dir1 : "input1\\ui\\ui_shop",
			dir2 : "input2\\UI-美術資源繁體\\28.商店",
			m : map[string]string{
				"ui_shop_module46_shangdian25.png":"module46_shangdiantubiao25.png",
				"ui_shop_module46_shangdian16.png":"module46_shangdiantubiao16.png",
				"ui_shop_module46_shangdian14.png":"module46_shangdiantubiao14.png",
				"ui_shop_module46_shangdian10.png":"module46_shangdiantubiao10.png",
				"ui_shop_module46_shangdian15.png":"module46_shangdiantubiao15.png",
				"ui_shop_module46_shangdian19.png":"module46_shangdiantubiao19.png",
			},
		},
		&fileprodata{
			dir1 : "input1\\ui\\ui_combat_finish",
			dir2 : "input2\\UI-美術資源繁體\\9.战斗结算",
			m : map[string]string{
				"ui_combat_finish_module03_jiesuan_zbqh_a.png":"module03_jiesuan_2.png",
				"ui_combat_finish_module03_jiesuan_wxhd_a.png":"module03_jiesuan_1.png",
				"ui_combat_finish_module03_jiesuan_zbqh_b.png":"module03_jiesuan_5.png",
				"ui_combat_finish_module03_jiesuan_wxsj_a.png":"module03_jiesuan_3.png",
				"ui_combat_finish_module03_jiesuan_shibai.png":"module03_jiesuan_shiabi.png",
				"ui_combat_finish_module03_jiesuan_wxsj_b.png":"module03_jiesuan_4.png",
				"ui_combat_finish_module03_jiesuan_wxhd_b.png":"module03_jiesuan_6.png",
			},
		},
		&fileprodata{
			dir1 : "input1\\ui\\ui_gain_items",
			dir2 : "input2\\UI-美術資源繁體\\0.共用（未完成）",
			m : map[string]string{
				"ui_gain_items_huode1.png":"huode.png",
			},
		},
		&fileprodata{
			dir1 : "input1\\ui\\ui_equip_tips",
			dir2 : "input2\\UI-美術資源繁體\\0.共用（未完成）",
			m : map[string]string{
				"ui_equip_tips_qianghuaanniu_a.png":"qianghua_a.png",
				"ui_equip_tips_qianghuaanniu_b.png":"qianghua_b.png",
			},
		},
		&fileprodata{
			dir1 : "input3\\img\\bg\\bg_common",
			dir2 : "input2\\UI-美術資源繁體\\0.共用（未完成）",
			m : map[string]string{
				"jiesuozhaoshi.png":"jiesuo.png",
			},
		},
		&fileprodata{
			dir1 : "input1\\ui\\ui_saodang",
			dir2 : "input2\\UI-美術資源繁體\\0.共用（未完成）",
			m : map[string]string{
				"ui_saodang_sdwc.png":"saodang.png",
			},
		},
		&fileprodata{
			dir1 : "input1\\ui\\ui_mapbuild",
			dir2 : "input2\\UI-美術資源繁體\\29.江湖事",
			m : map[string]string{
				"ui_mapbuild_module6_jianghushi_qiecha_a.png":"module6_jianghushi_qiecuo_a.png",
				"ui_mapbuild_module6_jianghushi_qiecha_b.png":"module6_jianghushi_qiecuo_b.png",
			},
		},
		&fileprodata{
			dir1 : "input1\\ui\\ui_major_role",
			dir2 : "input2\\UI-美術資源繁體\\26.混元",
			m : map[string]string{
				"ui_major_role_module98_hunyuan_9.png" :"2-1.png",
				"ui_major_role_module98_hunyuan_20.png":"2-2.png",
			},
		},
		&fileprodata{
			dir1 : "input1\\ui\\ui_major_role",
			dir2 : "input2\\UI-美術資源繁體\\7.时装",
			m : map[string]string{
				"ui_major_role_module11_rwsx_shizhuang.png" :"module40_shizhuang_15.png",
			},
		},
		&fileprodata{
			dir1 : "input1\\ui\\ui_meeting",
			dir2 : "input2\\UI-美術資源繁體\\31.武道之巅",
			m : map[string]string{
				"ui_meeting_module89_wdzd_24.png" :"adf.png",
			},
		},
		&fileprodata{
			dir1 : "input1\\ui\\ui_sign",
			dir2 : "input2\\UI-美術資源繁體\\27.活动\\1.签到\\签到",
			m : map[string]string{
				"ui_sign_module33_qiandao_buqian.png" :"補簽2.png",
			},
		},
		&fileprodata{
			dir1 : "input1\\ui\\ui_skill",
			dir2 : "input2\\UI-美術資源繁體\\24.武学",
			m : map[string]string{
				"ui_skill_module14_wuxue_anniu_xiulian.png" :"module14_wuxue_anniu _xiulian.png",
			},
		},
		&fileprodata{
			dir1 : "input1\\ui\\ui_combat",
			dir2 : "input2\\UI-美術資源繁體\\32.战斗",
			m : map[string]string{
				"ui_combat_module03_battle_anniu_shanbi.png" :"module03_battle_shanbi.png",
			},
		},
	}
)

type procret struct {
	Input1 string //简体资源
	Input2 string //繁体资源
	Output string //输出资源
}

type fileprodata struct {
	dir1 string
	dir2 string
	m map[string]string
}

//取繁体资源
/*
	&fileprodata{
		dir1 : "input3\\img\\bg\\bg_jianghu",
		dir2 : "input2\\UI-美術資源繁體\\11.江湖\\输出文件",
		m : map[string]string{
			"module6_jianghu_zhuluzhanchang_a.png":"module6_jianghu_xiakexing_a.png",
			"module6_jianghu_zhuluzhanchang_b.png":"module6_jianghu_xiakexing_b.png",
		},
	},
*/
func (f *fileprodata) proc() []*procret {
	ret := []*procret{}
	for k, v := range f.m {
		t := filepath.Join(f.dir1, k)
		d := createDir(t, inputDir1)
		nf := filepath.Join(d, k)
		of := filepath.Join(f.dir2, v)
		copy(of, nf)
		ret = append(ret, &procret{Input1: t, Input2 : of, Output: nf})
	}
	return ret
}

type dirprodata struct {
	dir1 string
	dir2 string
	convert func(string) (string, string)
}

//取交集 
/*
	&dirprodata{
		dir1 : "input3\\img\\bg\\bg_login",
		dir2 : "input2\\UI-美術資源繁體\\16.公告\\公告",
		//0.png -> 0.png 
		convert : func(str string) (string, string) {
			return str, str
		},
	},
*/
func (i *dirprodata) proc() []*procret {
	fs1 := getFileBase(i.dir1)
	fs2 := getFileBase(i.dir2)
	ret := []*procret{}
	for k, v := range fs2 {
		kc, nn := i.convert(k)
		if v1, ok := fs1[kc]; ok {
			d := createDir(v1, inputDir3)
			n := filepath.Join(d, nn)
			copy(v, n)
			ret = append(ret, &procret{Input1: v1, Input2 : v, Output: n})
		}
	}
	return ret
}

func getFileBase(path string) map[string]string {
	fileList := make(map[string]string)
    err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
        if strings.HasSuffix(path, ".png"){
			if _, ok := fileList[filepath.Base(path)]; ok{
				fmt.Println("double name, aaaaaaaaaaaaaaaaaaaaaaaaaaaa", path)
			}
			fileList[filepath.Base(path)] = path
        }
        return nil
    })

    if err != nil {
        fmt.Println(err)
		panic("err")
    }

	return fileList
}

func getFileList(path string) []string {
	fileList := []string{}
    err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
        if strings.HasSuffix(path, ".png"){
            fileList = append(fileList, path)
        }
        return nil
    })

    if err != nil {
        fmt.Println(err)
		panic("err")
    }

	return fileList
}

func checkDoubleName(files []string) bool {
	ret := true
	tmp := make(map[string][]string)
	for _, n := range files {
		b := filepath.Base(n)
		tmp[b] = append(tmp[b], n)
	}

	for k, v := range tmp {
		if len(v) > 1 {
			fmt.Println(k)
			for _, v1 := range v {
				fmt.Printf(" %v\n", v1)
			}
			ret = false
		}
	}

	return ret
}

func checkDoubleNameEx(files []string) bool {
	ret := true
	tmp := make(map[string][]string)
	for _, n := range files {
		//b := filepath.Base(n)
		b := delDirname(n)
		tmp[b] = append(tmp[b], n)
	}

	for k, v := range tmp {
		if len(v) > 1 {
			fmt.Println(k)
			for _, v1 := range v {
				fmt.Printf(" %v\n", v1)
			}
			ret = false
		}
	}

	return ret
}

// ui_login\ui_login_jian.png -> jian.png 
func delDirname(name string) string {
	//b: ui_login_jian.png
	b := filepath.Base(name)
	//l: ui_login 
	p := filepath.Dir(name)
	l := filepath.Base(p)
	i := strings.Index(b, l)
	if i  == -1 {
		//fmt.Println("format err.", name)
		return b
	}

	//return ui_login_jian.png - ui_login - 1(_)
	return b[len(l)+1:]
}

//path: input1\ui\ui_login\ui_login_jian.png 
//create dir: output\ui\ui_login\
//return: output\ui\ui_login\
func createDir(name, dir string) string{
	d := filepath.Dir(name)
	d = d[len(dir):]
	d = filepath.Join(outputDir, d)
	if _, err := os.Stat(d); os.IsNotExist(err) {
		os.MkdirAll(d, os.ModePerm)
	}
	return d
}

func copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

func RemoveContents(dir string) error {
    d, err := os.Open(dir)
    if err != nil {
        return err
    }
    defer d.Close()
    names, err := d.Readdirnames(-1)
    if err != nil {
        return err
    }
    for _, name := range names {
        err = os.RemoveAll(filepath.Join(dir, name))
        if err != nil {
            return err
        }
    }
    return nil
}

func main() {
	RemoveContents(outputDir)

	files1 := getFileList(inputDir1)
	files2 := getFileList(inputDir2)
	files3 := getFileList(inputDir3)

	fmt.Printf("input1:%d\n", len(files1))
	fmt.Printf("input2:%d\n", len(files2))
	fmt.Printf("input3:%d\n", len(files3))

	//重名检查
	//if !checkDoubleNameEx(files1) {
	if !checkDoubleName(files1) {
		fmt.Println("input1 double name")
		return
	}
	//if !checkDoubleName(files2) {
	if !checkDoubleName(files2) {
		fmt.Println("input2 double name")
		return
	}

	//繁体资源使用记录
	uf := make(map[string][]string)
	//key 短名字， value 路径
	s2 := make(map[string]string)
	for _, file := range files2 {
		s2[filepath.Base(file)] = file
		uf[file] = []string{}
	}

	//特定文件夹比较
	for _, v := range dirpro {
		list := v.proc()
		for _, v1 := range list {
			//记录已经使用的文件
			uf[v1.Input2] = append(uf[v1.Input2], v1.Input1)
		}
	}

	//特定文件比较
	for _, v := range filepro {
		list := v.proc()
		for _, v1 := range list {
			//记录已经使用的文件
			uf[v1.Input2] = append(uf[v1.Input2], v1.Input1)
		}
	}

	//操作input1和input2交集
	for _, file := range files1 {
		n1 := delDirname(file)
		if v, ok := s2[n1]; ok {
			_, notcheck := repeatlist[v]
			if !notcheck {
				//重复使用检测
				if v2, ok2 := uf[v]; ok2 && len(v2) >= 1 {
					//在已知名单中，不用提示错误
					if _, ok3 := repeatForbitlist[v]; !ok3 {
						fmt.Printf("check repeat, skip\n%v\n%v\n", file,  v)
						for _, v3 := range v2 {
							fmt.Printf("been use: %v \n", v3)
						}
					}
					continue
				}
			}

			d := createDir(file, inputDir1)
			n := filepath.Join(d, filepath.Base(file))
			_, err := copy(v, n)
			if err != nil {
				fmt.Println(err)
			}
			//fmt.Printf("%s\ncopy to\n%s\n", v, n)
			//记录已经使用的文件
			uf[v] = append(uf[v], file)
		}
	}

	fmt.Println("..............................")
	//显示使用情况
	fmt.Println("未使用的资源")
	for k, v := range uf {
		if _, ok := droplist[k]; !ok && len(v) == 0 {
			fmt.Println(k)
		}
	}
	fmt.Println("重复使用的资源")
	for k, v := range uf {
		if len(v) > 1 {
			fmt.Printf("count:%d, %v\n", len(v), k)
			for _, v1 := range v {
				fmt.Printf(" %v\n", v1)
			}
		}
	}

	//dorp
	fmt.Println("已经停止使用的资源", len(droplist))
	for k, _ := range droplist {
		fmt.Println(k)
	}

	//forbitlist
	fmt.Println("禁止在ui目录下重复使用的资源", len(repeatForbitlist))
	for k, _ := range repeatForbitlist {
		fmt.Println(k)
	}
}
