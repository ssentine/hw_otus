package hw03frequencyanalysis

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Change to true if needed.
var taskWithAsteriskIsCompleted = false

var text = `Как видите, он  спускается  по  лестнице  вслед  за  своим
	другом   Кристофером   Робином,   головой   вниз,  пересчитывая
	ступеньки собственным затылком:  бум-бум-бум.  Другого  способа
	сходить  с  лестницы  он  пока  не  знает.  Иногда ему, правда,
		кажется, что можно бы найти какой-то другой способ, если бы  он
	только   мог   на  минутку  перестать  бумкать  и  как  следует
	сосредоточиться. Но увы - сосредоточиться-то ему и некогда.
		Как бы то ни было, вот он уже спустился  и  готов  с  вами
	познакомиться.
	- Винни-Пух. Очень приятно!
		Вас,  вероятно,  удивляет, почему его так странно зовут, а
	если вы знаете английский, то вы удивитесь еще больше.
		Это необыкновенное имя подарил ему Кристофер  Робин.  Надо
	вам  сказать,  что  когда-то Кристофер Робин был знаком с одним
	лебедем на пруду, которого он звал Пухом. Для лебедя  это  было
	очень   подходящее  имя,  потому  что  если  ты  зовешь  лебедя
	громко: "Пу-ух! Пу-ух!"- а он  не  откликается,  то  ты  всегда
	можешь  сделать вид, что ты просто понарошку стрелял; а если ты
	звал его тихо, то все подумают, что ты  просто  подул  себе  на
	нос.  Лебедь  потом  куда-то делся, а имя осталось, и Кристофер
	Робин решил отдать его своему медвежонку, чтобы оно не  пропало
	зря.
		А  Винни - так звали самую лучшую, самую добрую медведицу
	в  зоологическом  саду,  которую  очень-очень  любил  Кристофер
	Робин.  А  она  очень-очень  любила  его. Ее ли назвали Винни в
	честь Пуха, или Пуха назвали в ее честь - теперь уже никто  не
	знает,  даже папа Кристофера Робина. Когда-то он знал, а теперь
	забыл.
		Словом, теперь мишку зовут Винни-Пух, и вы знаете почему.
		Иногда Винни-Пух любит вечерком во что-нибудь поиграть,  а
	иногда,  особенно  когда  папа  дома,  он больше любит тихонько
	посидеть у огня и послушать какую-нибудь интересную сказку.
		В этот вечер...`

var text1 = `Мороз и солнце; день чудесный!
	Еще ты дремлешь, друг прелестный —
	Пора, красавица, проснись:
	Открой сомкнуты негой взоры
	Навстречу северной Авроры,
	Звездою севера явись!
	Вечор, ты помнишь, вьюга злилась,
	На мутном небе мгла носилась;
	Луна, как бледное пятно,
	Сквозь тучи мрачные желтела,
	И ты печальная сидела —
	А нынче... погляди в окно:
	Под голубыми небесами
	Великолепными коврами,
	Блестя на солнце, снег лежит;
	Прозрачный лес один чернеет,
	И ель сквозь иней зеленеет,
	И речка подо льдом блестит.
	Вся комната янтарным блеском
	Озарена. Веселым треском
	Трещит затопленная печь.
	Приятно думать у лежанки.
	Но знаешь: не велеть ли в санки
	Кобылку бурую запречь?
	Скользя по утреннему снегу,
	Друг милый, предадимся бегу
	Нетерпеливого коня
	И навестим поля пустые,
	Леса, недавно столь густые,
	И берег, милый для меня.`

var text2 = `Why do you cry, Willy? 
			Why do you cry? 
			Why, Willy? 
			Why, Willy? 
			Why, Willy? Why?`

func TestTop10(t *testing.T) {
	t.Run("no words in empty string", func(t *testing.T) {
		require.Len(t, Top10(""), 0)
	})

	t.Run("positive test", func(t *testing.T) {
		if taskWithAsteriskIsCompleted {
			expected := []string{
				"а",         // 8
				"он",        // 8
				"и",         // 6
				"ты",        // 5
				"что",       // 5
				"в",         // 4
				"его",       // 4
				"если",      // 4
				"кристофер", // 4
				"не",        // 4
			}
			require.Equal(t, expected, Top10(text))
		} else {
			expected := []string{
				"он",        // 8
				"а",         // 6
				"и",         // 6
				"ты",        // 5
				"что",       // 5
				"-",         // 4
				"Кристофер", // 4
				"если",      // 4
				"не",        // 4
				"то",        // 4
			}
			require.Equal(t, expected, Top10(text))
		}
	})

	t.Run("More than 10 words test", func(t *testing.T) {
		if !taskWithAsteriskIsCompleted {
			expected := []string{
				"И",             // 5
				"ты",            // 3
				"в",             // 2
				"—",             // 2
				"А",             // 1
				"Авроры,",       // 1
				"Блестя",        // 1
				"Великолепными", // 1
				"Веселым",       // 1
				"Вечор,",        // 1
			}
			require.Equal(t, expected, Top10(text1))
		}
	})

	t.Run("Less than 10 words test", func(t *testing.T) {
		if !taskWithAsteriskIsCompleted {
			expected := []string{
				"Willy?", // 4
				"Why,",   // 3
				"Why",    // 2
				"do",     // 2
				"you",    // 2
				"Why?",   // 1
				"cry,",   // 1
				"cry?",   // 1
			}
			require.Equal(t, expected, Top10(text2))
		}
	})
}
