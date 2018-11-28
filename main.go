package main
import (
	"fmt"
	"os"
	"io"
	"io/ioutil"
	"strings"
	"path/filepath"
	"encoding/json"
)

var (
	dirpro = []*dirprodata {
		&dirprodata{
			dir1 : "icon\\jingjie",
			dir2 : "5.角色（人物属性）\\人物境界",
			//0.png -> jingjie_0.png 
			convert : func(str string) (fk, nn string) {
				fk = fmt.Sprintf("jingjie_%s", str)
				nn = fk
				return
			},
		},
		&dirprodata{
			dir1 : "bg\\bg_jianghu",
			dir2 : "11.江湖\\输出文件",
			//0.png -> 0.png 
			convert : func(str string) (string, string) {
				return str, str
			},
		},
		&dirprodata{
			dir1 : "bg\\bg_zlzc",
			dir2 : "30.逐鹿战场",
			//0.png -> 0.png 
			convert : func(str string) (string, string) {
				return str, str
			},
		},
		&dirprodata{
			dir1 : "build",
			dir2 : "3.主界面\\功能建筑（繁体）",
			//0.png -> 0.png 
			convert : func(str string) (string, string) {
				return str, str
			},
		},
		&dirprodata{
			dir1 : "bg\\bg_login",
			dir2 : "16.公告\\公告",
			//0.png -> 0.png 
			convert : func(str string) (string, string) {
				return str, str
			},
		},
		&dirprodata{
			dir1 : "bg\\bg_tower",
			dir2 : "33.修罗塔",
			//0.png -> 0.png 
			convert : func(str string) (string, string) {
				return str, str
			},
		},
	}

	filepro = []*fileprodata {
		&fileprodata{
			dir1 : "ui_major",
			dir2 : "11.江湖",
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
			dir1 : "bg\\bg_jianghu",
			dir2 : "11.江湖\\输出文件",
			m : map[string]string{
				"module6_jianghu_zhuluzhanchang_a.png":"module6_jianghu_xiakexing_a.png",
				"module6_jianghu_zhuluzhanchang_b.png":"module6_jianghu_xiakexing_b.png",
			},
		},
		&fileprodata{
			dir1 : "ui_lilian",
			dir2 : "10.历练秘境",
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
			dir1 : "ui_mijing",
			dir2 : "10.历练秘境",
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
			dir1 : "ui_shop",
			dir2 : "28.商店",
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
			dir1 : "ui_combat_finish",
			dir2 : "9.战斗结算",
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
			dir1 : "ui_gain_items",
			dir2 : "0.共用（未完成）",
			m : map[string]string{
				"ui_gain_items_huode1.png":"huode.png",
			},
		},
		&fileprodata{
			dir1 : "ui_equip_tips",
			dir2 : "0.共用（未完成）",
			m : map[string]string{
				"ui_equip_tips_qianghuaanniu_a.png":"qianghua_a.png",
				"ui_equip_tips_qianghuaanniu_b.png":"qianghua_b.png",
			},
		},
		&fileprodata{
			dir1 : "bg\\bg_common",
			dir2 : "0.共用（未完成）",
			m : map[string]string{
				"jiesuozhaoshi.png":"jiesuo.png",
			},
		},
		&fileprodata{
			dir1 : "ui_saodang",
			dir2 : "0.共用（未完成）",
			m : map[string]string{
				"ui_saodang_sdwc.png":"saodang.png",
			},
		},
		&fileprodata{
			dir1 : "ui_mapbuild",
			dir2 : "29.江湖事",
			m : map[string]string{
				"ui_mapbuild_module6_jianghushi_qiecha_a.png":"module6_jianghushi_qiecuo_a.png",
				"ui_mapbuild_module6_jianghushi_qiecha_b.png":"module6_jianghushi_qiecuo_b.png",
			},
		},
		&fileprodata{
			dir1 : "ui_major_role",
			dir2 : "26.混元",
			m : map[string]string{
				"ui_major_role_module98_hunyuan_9.png" :"2-1.png",
				"ui_major_role_module98_hunyuan_20.png":"2-2.png",
			},
		},
		&fileprodata{
			dir1 : "ui_major_role",
			dir2 : "7.时装",
			m : map[string]string{
				"ui_major_role_module11_rwsx_shizhuang.png" :"module40_shizhuang_15.png",
			},
		},
		&fileprodata{
			dir1 : "ui_meeting",
			dir2 : "31.武道之巅",
			m : map[string]string{
				"ui_meeting_module89_wdzd_24.png" :"adf.png",
			},
		},
		&fileprodata{
			dir1 : "ui_sign",
			dir2 : "27.活动\\1.签到\\签到",
			m : map[string]string{
				"ui_sign_module33_qiandao_buqian.png" :"補簽2.png",
			},
		},
		&fileprodata{
			dir1 : "ui_skill",
			dir2 : "24.武学",
			m : map[string]string{
				"ui_skill_module14_wuxue_anniu_xiulian.png" :"module14_wuxue_anniu _xiulian.png",
			},
		},
		&fileprodata{
			dir1 : "ui_combat",
			dir2 : "32.战斗",
			m : map[string]string{
				"ui_combat_module03_battle_anniu_shanbi.png" :"module03_battle_shanbi.png",
			},
		},
		&fileprodata{
			dir1 : "ui_story_info",
			dir2 : "0.共用（未完成）\\按钮",
			m : map[string]string{
				"ui_story_info_module_bangpai_62.png" : "module08__jqfb_17.png",
				"ui_story_info_module_bangpai_63.png" : "module08__jqfb_20.png",
				"ui_story_info_module6_jianghushi_jiejiao_a.png" : "module08__jqfb_18.png",
				"ui_story_info_module6_jianghushi_jiejiao_b.png" : "module08__jqfb_21.png",
				"ui_story_info_module09__tzjm_saodang03.png" :"module08__jqfb_9.png",
				"ui_story_info_module09__tzjm_saodang07.png" :"module08__jqfb_13.png",
				"ui_story_info_module09__tzjm_saodang08.png" :"module08__jqfb_10.png",
				"ui_story_info_module09__tzjm_saodang06.png" :"module08__jqfb_12.png",
				"ui_story_info_module09__tzjm_saodang04.png" : "module08__jqfb_16.png",
				"ui_story_info_module09__tzjm_saodang05.png" : "module08__jqfb_11.png",
				"ui_story_info_tzjm_tiaozhananniu_b.png" : "module08__jqfb_15.png",
				"ui_story_info_tzjm_tiaozhananniua.png" : "module08__jqfb_14.png",
			},
		},
		&fileprodata{
			dir1 : "ui_upgrade_gift",
			dir2 : "0.共用（未完成）",
			m : map[string]string{
				"ui_upgrade_gift_module33_zljl_6.png" : "module65_pata_7.png",
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

func (f *fileprodata) getDir(conf *configInfo) (dir1, dir2, dir3 string){
	if f.dir1[:3] == "ui_" {
		dir1 = filepath.Join(conf.Input1, f.dir1)
		dir3 = conf.Input1
	} else {
		dir1 = filepath.Join(conf.Input3, f.dir1)
		dir3 = conf.Input3
	}
	dir2 = filepath.Join(conf.Input2, f.dir2)
	return
}

//取繁体资源
/*
	&fileprodata{
		dir1 : "bg\\bg_jianghu",
		dir2 : "11.江湖\\输出文件",
		m : map[string]string{
			"module6_jianghu_zhuluzhanchang_a.png":"module6_jianghu_xiakexing_a.png",
			"module6_jianghu_zhuluzhanchang_b.png":"module6_jianghu_xiakexing_b.png",
		},
	},
*/
func (f *fileprodata) proc(conf *configInfo) []*procret {
	dir1, dir2, dir3 := f.getDir(conf)
	ret := []*procret{}
	for k, v := range f.m {
		t := filepath.Join(dir1, k)
		d := createDir(t, dir3, conf.Output)
		nf := filepath.Join(d, k)
		of := filepath.Join(dir2, v)
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

func (d *dirprodata) getDir(conf *configInfo) (dir1, dir2, dir3 string){
	if d.dir1[:3] == "ui_" {
		dir1 = filepath.Join(conf.Input1, d.dir1)
		dir3 = conf.Input1
	} else {
		dir1 = filepath.Join(conf.Input3, d.dir1)
		dir3 = conf.Input3
	}
	dir2 = filepath.Join(conf.Input2, d.dir2)
	return
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
func (i *dirprodata) proc(conf *configInfo) []*procret {
	dir1, dir2, dir3 := i.getDir(conf)
	fs1 := getFileBase(dir1)
	fs2 := getFileBase(dir2)
	ret := []*procret{}
	for k, v := range fs2 {
		kc, nn := i.convert(k)
		if v1, ok := fs1[kc]; ok {
			d := createDir(v1, dir3, conf.Output)
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

/*
name F:\tianxian_client\other\edit\ui\ui_zlzc\ui_zlzc_module83_zlzc_15.png
dir F:\tianxian_client\other\edit\ui
output C:\Users\Administrator\go\src\res_proc\output
return C:\Users\Administrator\go\src\res_proc\output\ui_zlzc
*/
func createDir(name, dir, output string) string{
	d := filepath.Dir(name)
	//F:\tianxian_client\other\edit\ui\ui_zlzc
	d2 := filepath.Dir(dir)
	//F:\tianxian_client\other\edit
	d = d[len(d2):]
	//\ui\ui_zlzc
	d = filepath.Join(output, d)
	//C:\Users\Administrator\go\src\res_proc\output\\ui\ui_zlzc
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

type configInfo struct {
	Input1 string `json:"ui_path"`
	Input2 string `json:"tradition_res"`
	Input3 string `json:"img_path"`
	Output string `json:"output_path"`
	DropFile []string `json:"drop_file"`
	RepeatFile []string `json:"repeat_file"`
	NotWarn []string `json:"not_warn"`
}

func loadConf() (*configInfo, error) {
	data, err := ioutil.ReadFile("conf.json")
    if err != nil {
        return nil, err
    }

	ci := &configInfo{}
    err = json.Unmarshal(data, &ci)
    if err != nil {
        return nil, err
    }
    return ci, nil
}

func getSepFile(files []string, dir string) map[string]bool {
	//"input2\\UI-美術資源繁體\\14.斗酒\\module33_doujiu_lose.png":true,
	ret := make(map[string]bool)
	for _, v := range files {
		str := filepath.Join(dir, v)
		ret[str] = true
	}
	return ret
}

func main() {
	conf, err := loadConf()
	if err != nil {
		fmt.Println(err)
		return
	}

	droplist := getSepFile(conf.DropFile, conf.Input2)
	repeatForbitlist := getSepFile(conf.NotWarn, conf.Input2)
	repeatlist := getSepFile(conf.RepeatFile, conf.Input2)

	RemoveContents(conf.Output)

	files1 := getFileList(conf.Input1)
	files2 := getFileList(conf.Input2)
	files3 := getFileList(conf.Input3)

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
		list := v.proc(conf)
		for _, v1 := range list {
			//记录已经使用的文件
			uf[v1.Input2] = append(uf[v1.Input2], v1.Input1)
		}
	}

	//特定文件比较
	for _, v := range filepro {
		list := v.proc(conf)
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

			d := createDir(file, conf.Input1, conf.Output)
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
