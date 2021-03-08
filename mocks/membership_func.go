package mocks

import "github.com/Orest-2/imfk/models"

func mockMembershipFunc() {
	obj := []models.MembershipFunc{
		{
			Name:      "Трикутна функції належності",
			Code:      "trimf",
			Type:      "2d",
			ParamsCnt: 3,
		},
		{
			Name:      "Трапециідальна функція належності",
			Code:      "trapmf",
			Type:      "2d",
			ParamsCnt: 4,
		},
		{
			Name:      "Квадратичний Z-сплайн",
			Code:      "szmf",
			Type:      "2d",
			ParamsCnt: 2,
		},
		{
			Name:      "Гармонійний Z-сплайн",
			Code:      "hzmf",
			Type:      "2d",
			ParamsCnt: 2,
		},
		{
			Name:      "Z-сигмоїдальна функція належності",
			Code:      "zsigmf",
			Type:      "2d",
			ParamsCnt: 2,
		},
		{
			Name:      "Z-лінійна функція належності",
			Code:      "zlinemf",
			Type:      "2d",
			ParamsCnt: 2,
		},
		{
			Name:      "Квадратичний S-сплайн",
			Code:      "ssmf",
			Type:      "2d",
			ParamsCnt: 2,
		},
		{
			Name:      "Гармонійний S-сплайн",
			Code:      "hsmf",
			Type:      "2d",
			ParamsCnt: 2,
		},
		{
			Name:      "S-сигмоїдальна функція належності",
			Code:      "ssigmf",
			Type:      "2d",
			ParamsCnt: 2,
		},
		{
			Name:      "S-лінійна функція належності",
			Code:      "slinemf",
			Type:      "2d",
			ParamsCnt: 2,
		},
		{
			Name:      "Дзвоноподібна функція належності",
			Code:      "gbellmf",
			Type:      "2d",
			ParamsCnt: 3,
		},
		{
			Name:      "Гаусова функція належності",
			Code:      "gaussmf",
			Type:      "2d",
			ParamsCnt: 2,
		},
		{
			Name:      "Конусоподібна функція належності",
			Code:      "conemf",
			Type:      "3d",
			ParamsCnt: 2,
		},
		{
			Name:      "Пірамідальна функція належності",
			Code:      "pyrammf",
			Type:      "3d",
			ParamsCnt: 2,
		},
		{
			Name:      "Трапецієподібно-пірамідальна функція належності",
			Code:      "trappyrammf",
			Type:      "3d",
			ParamsCnt: 3,
		},
		{
			Name:      "Узагальнена сигмоїдальна функція належності",
			Code:      "gsigmf",
			Type:      "3d",
			ParamsCnt: 2,
		},
		{
			Name:      "Дзвоноподібна функція належності",
			Code:      "gbell3dmf",
			Type:      "3d",
			ParamsCnt: 3,
		},
		{
			Name:      "Гаусова функція належності",
			Code:      "gauss3dmf",
			Type:      "3d",
			ParamsCnt: 2,
		},
		{
			Name:      "Гіперболоїдна функція належності",
			Code:      "hyperbolmf",
			Type:      "3d",
			ParamsCnt: 2,
		},
		{
			Name:      "Еліпсоїдна функція належності",
			Code:      "ellipsmf",
			Type:      "3d",
			ParamsCnt: 2,
		},
	}

	for i, element := range obj {
		element.BaseModel.ID = i + 1
		models.DB.Save(&element)
	}
}
