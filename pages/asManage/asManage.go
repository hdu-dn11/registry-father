package asManage

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"net"
	"registry-father/model"
	"registry-father/services"
	"strconv"
	"strings"
)

func Load(app *tview.Application) tview.Primitive {
	asList := tview.NewList().ShowSecondaryText(false)
	asList.SetBorder(true).SetTitle("AS list")

	//asInfo := tview.NewForm()
	//asInfo.SetBorder(true).SetTitle("AS info")

	asInfo := NewASInfoForm(app, asList)

	asList.SetFocusFunc(func() {
		asList.Clear()

		asList.AddItem("Create AS", "", 0, func() {
			asInfo.Load(&model.ASInfo{})
			app.SetFocus(asInfo)
		})
		asInfo.Load(&model.ASInfo{})

		asInfoList, err := services.GetASInfoList()
		if err != nil {
			panic(err)
		}
		for _, info := range asInfoList {
			data := info
			asList.AddItem(fmt.Sprintf("AS%d", info.ASN), "", 0, func() {
				asInfo.Load(data)
				app.SetFocus(asInfo)
			})
		}

		asList.SetChangedFunc(func(index int, mainText string, secondaryText string, shortcut rune) {
			if asList.GetCurrentItem() > 0 {
				idx := asList.GetCurrentItem()
				asInfo.Load(asInfoList[idx-1])
			} else {
				asInfo.Load(&model.ASInfo{})
			}
		})
	})

	//asList.SetSelectedFunc(func(i int, s string, s2 string, r rune) {
	//	idx := asList.GetCurrentItem()
	//	asInfo.Load(asInfoList[idx+1])
	//	app.SetFocus(asInfo)
	//})

	return tview.NewFlex().
		AddItem(asList, 0, 1, true).
		AddItem(asInfo, 0, 2, true)
}

type ASInfoForm struct {
	*tview.Form
	defaultBgColor tcell.Color
	app            *tview.Application
	parent         tview.Primitive
}

// NewASInfoForm app is the application and if parent is not nil then if form finished,
// it will focus to parent
func NewASInfoForm(app *tview.Application, parent tview.Primitive) *ASInfoForm {
	form := tview.NewForm()
	form.SetBorder(true).SetTitle("AS info")
	form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyESC {
			app.SetFocus(parent)
		}
		return event
	})
	return &ASInfoForm{
		Form:           form,
		defaultBgColor: form.GetBackgroundColor(),
		app:            app,
		parent:         parent,
	}
}

func (f *ASInfoForm) Load(info *model.ASInfo) {
	tmpInfo := info.Clone()

	f.SetBackgroundColor(f.defaultBgColor)
	f.Clear(true)

	f.AddInputField("ASN", strconv.Itoa(int(tmpInfo.ASN)), 0,
		func(textToCheck string, lastChar rune) bool {
			parseASN, err := strconv.ParseUint(textToCheck, 10, 32)
			if err != nil {
				return false
			}
			tmpInfo.ASN = uint32(parseASN)
			return true
			//return services.InDN11(net.ParseIP(textToCheck))
		}, nil)

	f.AddInputField("Owner", tmpInfo.Owner, 0, nil, func(text string) {
		tmpInfo.Owner = strings.Trim(text, " ")
	})

	ipv4checker := func(cidr string) error {
		ip, ipNet, err := net.ParseCIDR(cidr)
		if err != nil {
			return fmt.Errorf("failed to parse")
		}
		if !services.InDN11(*ipNet) {
			return fmt.Errorf("%s is not in DN11", ip)
		}
		return nil
	}

	IPv4Input := tview.NewTextArea()
	IPv4Input.SetLabel("IPv4").
		SetText(strings.Join(tmpInfo.IPv4, "\n"), true).
		SetSize(len(tmpInfo.IPv4)+1, len("255.255.255.255/32")).
		SetChangedFunc(func() {
			input := IPv4Input.GetText()
			tmpInfo.IPv4 = strings.Split(input, "\n")
			IPv4Input.SetSize(len(tmpInfo.IPv4)+1, len("255.255.255.255/32"))
			valid := true
			for _, s := range tmpInfo.IPv4 {
				if ipv4checker(s) != nil {
					valid = false
					break
				}
			}
			if !valid {
				IPv4Input.SetLabelStyle(tcell.StyleDefault.Foreground(tcell.ColorRed))
				IPv4Input.SetTextStyle(tcell.StyleDefault.Background(tcell.ColorRed).Foreground(tview.Styles.TertiaryTextColor))
				f.SetBackgroundColor(tcell.ColorRed)
			} else {
				f.SetBackgroundColor(f.defaultBgColor)
			}
		})
	f.AddFormItem(IPv4Input)

	IPv6Input := tview.NewTextArea()
	IPv6Input.SetLabel("IPv6").
		SetText(strings.Join(tmpInfo.IPv6, "\n"), true).
		SetSize(len(tmpInfo.IPv6)+1, len("1111:1111:1111:1111:1111:1111:1111:1111/128")).
		SetChangedFunc(func() {
			input := IPv6Input.GetText()
			tmpInfo.IPv6 = strings.Split(input, "\n")
			IPv6Input.SetSize(len(tmpInfo.IPv6)+1, len("1111:1111:1111:1111:1111:1111:1111:1111/128"))
		})
	f.AddFormItem(IPv6Input)

	f.AddButton("SAVE", func() {
		_ = services.SaveASInfo(&tmpInfo)
		f.app.SetFocus(f.parent)

	})

	f.AddButton("CANCEL", func() {
		f.app.SetFocus(f.parent)
	})

	f.AddButton("DELETE", func() {
		_ = services.DeleteASInfo(&tmpInfo)
		f.app.SetFocus(f.parent)
	})
}
