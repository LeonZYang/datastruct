package bisect

func Bisect(arr []int, target int) int {
	num := len(arr)
	low, high := 0, num-1

	for low <= high {
		mid := low + (high-low)/2
		if arr[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	if high >= 0 && high < num && arr[high] == target {
		return high
	}
	return -1
}

func BisectLeft(arr []int, target int) int {
	num := len(arr)
	low, high := 0, num-1

	for low <= high {
		mid := low + (high-low)/2
		if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if low >= 0 && low < num && arr[low] == target {
		return low
	}
	return -1
}

func BisectRight(arr []int, target int) int {
	num := len(arr)
	low, high := 0, num-1

	for low <= high {
		mid := low + (high-low)/2
		if arr[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	if high >= 0 && high < num && arr[high] == target {
		return high + 1
	}
	return -1
}
